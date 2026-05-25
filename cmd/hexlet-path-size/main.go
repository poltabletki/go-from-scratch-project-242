package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of file or directory",
		ArgsUsage: "<path>",
		Action: func(_ context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				return fmt.Errorf("path is required")
			}

			path := c.Args().Get(0)
			size, err := code.GetPathSize(path, false, false, false)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
