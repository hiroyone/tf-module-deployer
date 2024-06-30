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

	// Create .tf-module directory if it doesn't exist
	tfModuleDir := filepath.Join(cwd, ".tf-module")
	if _, err := os.Stat(tfModuleDir); os.IsNotExist(err) {
		if err := os.Mkdir(tfModuleDir, 0755); err != nil {
			handleError("Error creating .tf-module directory", err)
		}
	}

	// Changing directory to .tf-module directory
	if err := os.Chdir(tfModuleDir); err != nil {
		handleError("Error changing directory to .tf-module", err)
	}

	fmt.Println("Changed directory to .tf-module")

	// Move main.tf file to .tf-module directory
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
