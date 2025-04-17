package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "sbx",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "allow-file",
				Usage: "path to the allow file",
			},
			&cli.StringFlag{
				Name:  "deny-file",
				Usage: "path to the deny file",
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
