package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"tf-module-deployer/utils"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the Terraform module",
	Run: func(cmd *cobra.Command, args []string) {
		deployCommand()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}

func deployCommand() {
	fmt.Println("Executing deploy command...")

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		utils.HandleError("Error getting current directory", err)
	}

	tfModuleDir := filepath.Join(cwd, ".tf-module")

	// Create a deployment directory if it doesn't exist
	if err := utils.CreateOrOverwriteDirectory(tfModuleDir); err != nil {
		fmt.Printf("Error creating or overwriting %s directory: %v\n", tfModuleDir, err)
		os.Exit(1)
	}

	// Changing directory to the created directory
	if err := os.Chdir(tfModuleDir); err != nil {
		utils.HandleError(
			fmt.Sprintf("Error changing directory to %s", tfModuleDir), err)
	}
	fmt.Println("Changed directory to .tf-module")

	// Copying main.tf file to .tf-module directory
	srcFile := filepath.Join(cwd, "main.tf")
	dstFile := filepath.Join(tfModuleDir, "main.tf")
	if err := utils.CopyFile(srcFile, dstFile); err != nil {
		utils.HandleError("Error moving main.tf file", err)
	}
	fmt.Printf("Moved main.tf to %s directory", tfModuleDir)
}

