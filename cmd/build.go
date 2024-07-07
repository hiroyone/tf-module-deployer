package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"tf-module-deployer/config"
	"tf-module-deployer/utils"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the Terraform module",
	Long:  fmt.Sprintf(`This command moves the %s file to the %s directory.`, config.FileName, config.ModuleDir),
	Run: func(cmd *cobra.Command, args []string) {
		buildCommand()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func buildCommand() {
	fmt.Println("Executing deploy command...")

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		utils.HandleError("Error getting current directory", err)
	}
	tfModuleDir := filepath.Join(cwd, config.ModuleDir)

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
	fmt.Printf("Changed directory to %s", tfModuleDir)

	// Copying main.tf file to .tf-module directory
	srcFile := filepath.Join(cwd, config.FileName)
	dstFile := filepath.Join(tfModuleDir, config.FileName)
	if err := utils.CopyFile(srcFile, dstFile); err != nil {
		utils.HandleError(
			fmt.Sprintf("Error moving %s file", config.FileName), err)
	} else {
		fmt.Printf("Moved %s to %s directory", config.FileName, tfModuleDir)
	}
}
