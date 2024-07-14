package cmd

import (
	"tf-module-deployer/config"
	"tf-module-deployer/utils"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Run terraform plan",
	Long:  `This command runs terraform plan in the .tf-module-deployer/examples directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		planCommand()
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}

func planCommand() {
	// Run terraform plan
	err := utils.RunTerraformCommand("plan", config.ModuleDir)
	utils.HandleErrorfLn("Failed to run terraform plan in %s", err, config.ModuleDir)
}
