package main

import (
	"os"

	"github.com/spf13/cobra"

	"my-project/cmd/app/start"
)

const example = `go run cmd/app/main.go start -e env.yaml`

func main() {
	root := cobra.Command{
		Use:     "app",
		Short:   "项目入口执行文件",
		Example: example,
	}

	root.PersistentFlags().StringP("env", "e", "env.yaml", "环境文件")

	root.AddCommand(start.NewCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
