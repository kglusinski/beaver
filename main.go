package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"

	_ "embed"
)

//go:embed templates/main.template.go
var mainTemplate string

func main() {
	app := &cli.App{
		Name:        "beaver",
		HelpName:    "",
		Usage:       "",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "v1.0.0",
		Description: "Beaver CLI tool helps Golang devs to quickly setup new projects",
		Commands: []*cli.Command{
			{
				Name:        "newProject",
				Usage:       "beaver newProject [--name project_name]",
				Description: "it generates newProject project structure",
				Action:      newProject,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "--name <project_name>",
						Required: true,
					},
				},
			},
		},
		Compiled:               time.Time{},
		Authors:                []*cli.Author{
			{
				Name:  "Kamil Głusiński",
				Email: "kontakt@zaprogramowani.dev",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newProject(c *cli.Context) error {
	projectName := c.String("name")

	err := os.Mkdir(projectName, 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/main.go", projectName), []byte(mainTemplate), 0644)
	if err != nil {
		log.Printf("cannot create main file, err: %v", err)
		return err
	}

	return nil
}
