package utils

import (
	"fmt"
	"os"
	"os/exec"
)
func RunTerraformCommand(command, dir string) error {

    // Check if terraform init is needed
    if _, err := os.Stat(fmt.Sprintf("%s/.terraform", dir)); os.IsNotExist(err) {
        fmt.Println("Running terraform init...")
        err := runCommand("init", dir)
        if err != nil {
            return err
        }
    }

    // Run the specified terraform command
    return runCommand(command, dir)
}

func runCommand(command, dir string) error {
    cmd := exec.Command("sh", "-c", fmt.Sprintf("terraform %s", command))
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}