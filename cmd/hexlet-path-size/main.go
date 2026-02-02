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
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Value:   false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			if path == "" {
				return fmt.Errorf("path is required")
			}

			human := cmd.Bool("human")

			return getPathSize(ctx, path, human)
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func getPathSize(ctx context.Context, path string, human bool) error {
	size, err := code.GetPathSize(path, human)
	if err != nil {
		return err
	}

	fmt.Println(size)

	return nil
}
