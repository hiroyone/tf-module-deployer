package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tf-module-deployer <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "deploy":
		deployCommand()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func deployCommand() {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Create .tf-module directory if it doesn't exist
	tfModuleDir := filepath.Join(cwd, ".tf-module")
	if _, err := os.Stat(tfModuleDir); os.IsNotExist(err) {
		if err := os.Mkdir(tfModuleDir, 0755); err != nil {
			fmt.Printf("Error creating .tf-module directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Move current directory to .tf-module directory
	if err := os.Chdir(tfModuleDir); err != nil {
		fmt.Printf("Error changing directory to .tf-module: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Moved current directory to .tf-module")

	// Move main.tf file to .tf-module directory
	srcFile := filepath.Join(cwd, "main.tf")
	dstFile := filepath.Join(tfModuleDir, "main.tf")

	if err := moveFile(srcFile, dstFile); err != nil {
		fmt.Printf("Error moving main.tf file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Moved main.tf to .tf-module directory")
}

func moveFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// Remove original file after successful copy
	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}
