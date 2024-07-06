package cmd

import (
	"fmt"
	"tf-module-deployer/config"
	"tf-module-deployer/utils"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
    Use:   "apply",
    Short: "Run terraform apply",
    Long:  `This command runs terraform apply in the .tf-module-deployer/examples directory.`,
    Run: func(cmd *cobra.Command, args []string) {
        err := utils.RunTerraformCommand("apply", config.ModuleDir)
        utils.HandleError(
			fmt.Sprintf("Failed to run terraform apply in %s", config.ModuleDir),err)
    },
}

func init() {
    rootCmd.AddCommand(applyCmd)
}