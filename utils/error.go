package utils

import (
	"os"
)

// HandleErrorfLn is a utility function that handles errors and prints a formatted error message.
// If an error is not nil, it prints the formatted error message along with the error and exits the program with status code 1.
func HandleErrorfLn(format string, err error, a ...interface{}) {
	if err != nil {
		PrintfLn(format+": %v", append(a, err)...)
		os.Exit(1)
	}
}
