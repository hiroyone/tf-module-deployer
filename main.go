// main.go
package main

import (
	"fmt"
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
}
