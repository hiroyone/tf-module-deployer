package utils

import "fmt"

// PrintfLn formats according to a format specifier and writes to standard output, followed by a newline.
func PrintfLn(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}
