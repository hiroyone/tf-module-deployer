package main

import (
	"fmt"
	"io"
	"os"
)

// createOrOverwriteDirectory creates a directory with the specified folderName in the current working directory (cwd).
// If the directory already exists, it deletes all existing contents and recreates it.
// It returns the absolute path to the newly created or overwritten directory.
func createOrOverwriteDirectory(dirPath string) error {
	// Remove directory if it exists
	if _, err := os.Stat(dirPath); err == nil {
		if err := os.RemoveAll(dirPath); err != nil {
			return fmt.Errorf("failed to remove existing directory %s: %w", dirPath, err)
		}
	}

	// Create new directory
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	return nil
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("error copying file contents: %w", err)
	}

	return nil
}
