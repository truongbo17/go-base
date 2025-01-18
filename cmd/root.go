package cmd

import (
	"embed"
	"fmt"
	"go-base/cmd/cli"
	"go-base/cmd/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ggb",
	Short: "Go-Gin-Base quickly build and develop web applications. restful API, microservice...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Go-Gin-Base. Use -h to see more commands")
	},
}

func init() {
	rootCmd.AddCommand(server.StartServerCmd)
	rootCmd.AddCommand(cli.VersionCmd)
}

func Execute(fs embed.FS) {
	var _ embed.FS

	_ = fs

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
