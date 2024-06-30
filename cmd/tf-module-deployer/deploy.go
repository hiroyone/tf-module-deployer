package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func deployCommand() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		handleError("Error getting current directory", err)
	}

	tfModuleDir := filepath.Join(cwd, ".tf-module")

	// Create .tf-module directory if it doesn't exist
	if err := createOrOverwriteDirectory(tfModuleDir); err != nil {
		fmt.Printf("Error creating or overwriting .tf-module directory: %v\n", err)
		os.Exit(1)
	}

	// Changing directory to .tf-module directory
	if err := os.Chdir(tfModuleDir); err != nil {
		handleError("Error changing directory to .tf-module", err)
	}
	fmt.Println("Changed directory to .tf-module")

	// Copying main.tf file to .tf-module directory
	srcFile := filepath.Join(cwd, "main.tf")
	dstFile := filepath.Join(tfModuleDir, "main.tf")
	if err := copyFile(srcFile, dstFile); err != nil {
		handleError("Error moving main.tf file", err)
	}
	fmt.Println("Moved main.tf to .tf-module directory")
}

func handleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}
