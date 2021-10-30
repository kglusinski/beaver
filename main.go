package main

import (
	"beaver/generate"
	_ "embed"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
		Email: "kontakt@zaprogramowani.dev",
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
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
