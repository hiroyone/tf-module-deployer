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

	moduleDir := ".tf-module"
	tfModuleDir := filepath.Join(cwd, moduleDir)

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
	fileName := "main.tf"
	srcFile := filepath.Join(cwd, fileName)
	dstFile := filepath.Join(tfModuleDir, fileName)
	if err := utils.CopyFile(srcFile, dstFile); err != nil {
		utils.HandleError(
			fmt.Sprintf("Error moving %s file", fileName), err)
	} else {	
		fmt.Printf("Moved %s to %s directory", fileName, tfModuleDir)
	}

       // Run terraform plan
	   err = utils.RunTerraformPlan(tfModuleDir)
	   utils.HandleError(
		fmt.Sprintf("Failed to run terraform plan in %s", tfModuleDir), err)

}

