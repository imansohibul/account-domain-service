package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "account-domain-service",
		Usage: "Webservice for accounts",
		Authors: []*cli.Author{
			{
				Name:  "Mochamad Sohibul Iman",
				Email: "iman@imansohibul.my.id",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{"a"},
				Usage:   "Start account service API server",
				Action:  RestAPI,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "address",
						Value: ":8080",
						Usage: "The address parameter defines the server address and port number that the server will listen on.",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
