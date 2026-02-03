package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
				Value:   false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			if path == "" {
				return fmt.Errorf("path is required")
			}

			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")
			all := cmd.Bool("all")

			return getPathSize(path, recursive, human, all)
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func getPathSize(path string, recursive, human, all bool) error {
	size, err := code.GetPathSize(path, recursive, human, all)
	if err != nil {
		return err
	}

	fmt.Printf("%s\t%s", size, path)

	return nil
}
