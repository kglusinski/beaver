package main

import (
	_ "embed"
	"log"
	"os"

	"beaver/generate"
	"github.com/urfave/cli/v2"
)

const (
	CommandName = "beaver"
)

func main() {
	app := cli.NewApp()
	app.Name = CommandName
	app.Version = "v1.0.0"
	app.Description = "Beaver CLI tool helps Golang devs to quickly setup new projects"
	app.Usage = ""
	app.Authors = []*cli.Author{{
		Name:  "Kamil Głusiński",
		Email: "kontakt@inzkawka.pl",
	}}
	app.Commands = []*cli.Command{
		{
			Name:        "new",
			Usage:       "beaver new [--name project_name]",
			Description: "it generates new project structure",
			Action:      generate.NewProject,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "--name project_name",
					Required: false,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
