package utils

import (
	"os"
	"os/exec"
)
func RunTerraformPlan(dir string) error {
    cmd := exec.Command("terraform", "plan")
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}