package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tf-module-deployer",
	Short: "CLI tool for deploying Terraform modules",
	Run: func(cmd *cobra.Command, args []string) {
		// Default action if no subcommand is specified
		cmd.Usage()
	},
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        // Exit the program if there's an error
        os.Exit(1)
    }
}