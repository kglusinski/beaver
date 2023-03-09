package template

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const defaultFilePerm = 0644

type StandardTemplate struct {
	Name     string
	template embed.FS
}

func NewStandardTemplate(name string, template embed.FS) *StandardTemplate {
	return &StandardTemplate{
		Name:     name,
		template: template,
	}
}

func (s *StandardTemplate) Generate() error {
	err := createDirTree(s.Name)
	if err != nil {
		log.Printf("cannot create project tree, err: %v", err)
		return err
	}

	mainTemplate, err := s.template.ReadFile("standard/main.go.template")
	if err != nil {
		log.Printf("cannot open main.go template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/cmd/main/main.go", s.Name), mainTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create main file, err: %v", err)
		return err
	}

	configTemplate, err := s.template.ReadFile("standard/config.go.template")
	if err != nil {
		log.Printf("cannot open config.go template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/cmd/main/config.go", s.Name), configTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create config file, err: %v", err)
		return err
	}

	makeTemplate, err := s.template.ReadFile("standard/Makefile.template")
	if err != nil {
		log.Printf("cannot open Makefile template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/Makefile", s.Name), makeTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create Makefile, err: %v", err)
		return err
	}

	gomodTemplate, err := s.template.ReadFile("standard/go.mod.template")
	if err != nil {
		log.Printf("cannot open Makefile template, err: %v", err)
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/go.mod", s.Name), gomodTemplate, defaultFilePerm)
	if err != nil {
		log.Printf("cannot create go.mod, err: %v", err)
		return err
	}

	replacement := bytes.Replace(gomodTemplate, []byte("_PROJECT_NAME_"), []byte(s.Name), -1)
	if err = ioutil.WriteFile(fmt.Sprintf("%s/go.mod", s.Name), replacement, defaultFilePerm); err != nil {
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
