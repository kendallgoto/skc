//go:build !linux
// +build !linux

package execute

func Run(code string) (string, error) {
	return "# Execution is not supported on this build", nil
}
