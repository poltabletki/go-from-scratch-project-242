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
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
				DefaultText: "false",
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				return fmt.Errorf("path is required")
			}

			path := c.Args().Get(0)
			recursive := c.Bool("recursive")
			human := c.Bool("human")
			all := c.Bool("all")
			size, err := code.GetPathSize(path, recursive, human, all)
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
