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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			if path == "" {
				return fmt.Errorf("path is required")
			}

			return pathSize(ctx, path)
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func pathSize(ctx context.Context, path string) error {
	size, err := code.GetSize(path)
	if err != nil {
		return err
	}

	fmt.Println(size)

	return nil
}
