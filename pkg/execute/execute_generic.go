package execute

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/uuid"
)

func Run(code string) (string, error) {
	//return "Execution is not supported on this machine", nil
	executionId := uuid.New().String()
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	sandboxPath := path.Join(cwd, "execution", "sandbox")
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}
	containerConfig := &container.Config{
		Image: "skc-sandbox",
		Cmd:   []string{"/app/sandbox/" + executionId + ".py"},
	}
	hostConfig := &container.HostConfig{
		Runtime:    "runsc",
		AutoRemove: true,
		Binds:      []string{sandboxPath + ":/app/sandbox"},
	}

	resp, err := cli.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		nil,
		nil,
		executionId,
	)
	if err != nil {
		return "", err
	}
	// save python
	err = os.WriteFile(path.Join(sandboxPath, executionId+".py"), []byte(code), 0666)
	defer os.Remove(path.Join(sandboxPath, executionId+".py"))
	if err != nil {
		return "", err
	}
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}
	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", err
		}
	case <-statusCh:
	}
	fmt.Printf("finished %s\n", resp.ID)
	containerLogs, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", err
	}
	stdOut := new(bytes.Buffer)
	stdErr := new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdOut, stdErr, containerLogs)
	if err != nil {
		return "", err
	}
	// fmt.Printf("stdout: \"%s\"\n", stdOut)
	// fmt.Printf("stderr: \"%s\"\n", stdErr)
	return stdOut.String(), nil
}
