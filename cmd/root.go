package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tf-module-deployer",
	Short: "A CLI to deploy Terraform modules",
	Long:  `tf-module-deployer is a CLI that helps in deploying Terraform modules to a server.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
