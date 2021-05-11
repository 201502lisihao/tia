package start

import (
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"my-project/inject"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "启动项目",
		RunE:  run,
	}
}

func run(cmd *cobra.Command, _ []string) error {
	envPath, err := cmd.Parent().PersistentFlags().GetString("env")
	if err != nil {
		return err
	}

	app, cleanup, err := inject.InitApp(envPath)
	if err != nil {
		return err
	}

	defer cleanup()

	var g errgroup.Group
	g.Go(func() error {
		return app.Enter.BootGrpcServer()
	})

	g.Go(func() error {
		return app.Enter.BootHttpServer()
	})

	return g.Wait()
}
