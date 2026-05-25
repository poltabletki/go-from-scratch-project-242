package main

import (
	"code"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "path is required")
		printUsage(os.Stderr)
		os.Exit(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage(os.Stdout)
		return
	}

	path := os.Args[1]
	size, err := code.GetPathSize(path, false, false, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\t%s\n", size, path)
}

func printUsage(out *os.File) {
	fmt.Fprintln(out, "Usage: hexlet-path-size <path>")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Options:")
	fmt.Fprintln(out, "  -h, --help    show this help message")
}
