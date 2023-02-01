package main

import (
	"github.com/kendallgoto/skc/pkg/server/skcserver"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	defer sugar.Sync()

	var root = &cobra.Command{
		Use:           "skc-server",
		Short:         "SKC Processing Server",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// load and DI any config here
			innerLogger, err := zap.NewDevelopment()
			if err != nil {
				return err
			}
			innerSugar := innerLogger.Sugar()
			defer innerSugar.Sync()

			innerSugar.Infof("Starting SKC server at %s\n", ":8080")

			return skcserver.Start(innerSugar)
		},
	}
	err := root.Execute()
	if err != nil {
		sugar.Fatalln(err.Error())
	}
}
