package generate

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

const defaultFilePerm = 0644

//go:embed templates/*
var templates embed.FS

func NewProject(c *cli.Context) error {
	projectName := c.String("name")

	err := createDirTree(projectName)
	if err != nil {
		log.Printf("cannot create project tree, err: %v", err)
		return err
	}

	mainTemplate, err := templates.ReadFile("templates/main.go.template")
	if err != nil {
		log.Printf("cannot open main.go template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/cmd/main/main.go", projectName), mainTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create main file, err: %v", err)
		return err
	}

	makeTemplate, err := templates.ReadFile("templates/Makefile.template")
	if err != nil {
		log.Printf("cannot open Makefile template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/Makefile", projectName), makeTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create Makefile, err: %v", err)
		return err
	}

	gomodTemplate, err := templates.ReadFile("templates/go.mod.template")
	if err != nil {
		log.Printf("cannot open Makefile template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/go.mod", projectName), gomodTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create go.mod, err: %v", err)
		return err
	}

	replacement := bytes.Replace(gomodTemplate, []byte("_PROJECT_NAME_"), []byte(projectName), -1)
	if err = ioutil.WriteFile(fmt.Sprintf("%s/go.mod", projectName), replacement, defaultFilePerm); err != nil {
		log.Printf("cannot create go.mod, err: %v", err)
		return err
	}

	return nil
}

func createDirTree(projectName string) error {
	err := os.Mkdir(projectName, 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/cmd", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/cmd/main", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/internal", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/pkg", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/testdata", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/tests", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	err = os.Mkdir(fmt.Sprintf("%s/docs", projectName), 0700)
	if err != nil {
		log.Printf("cannot create directory, err: %v", err)
		return err
	}

	return nil
}
