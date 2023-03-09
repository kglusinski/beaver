package generate

import (
	"errors"
	"fmt"
	"log"
	"os/exec"

	template2 "beaver/generate/template"
	templateDir "beaver/templates"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

const (
	FlagName      = "name"
	FlagStructure = "structure"
)

var templates = templateDir.Templates

func NewProject(c *cli.Context) error {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("name cannot be empty")
		}
		return nil
	}

	promptName := promptui.Prompt{
		Label:    "Project name",
		Default:  "",
		Validate: validate,
	}

	projectName, err := promptName.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
	}

	promptType := promptui.Select{
		Label: "Select project structure",
		Items: []string{"standard", "hexagonal"},
	}

	_, projectStructure, err := promptType.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v", err)
	}

	switch projectStructure {
	case "standard":
		template := template2.NewStandardTemplate(projectName, templates)
		template.Generate()
	case "hexagonal":
		template := template2.NewHexagonalTemplate(projectName, templates)
		template.Generate()
	}

	cmd := exec.Command("go", "mod", "download")
	err = cmd.Run()
	if err != nil {
		log.Printf("cannot download dependencies, err: %v", err)
		return err
	}

	return nil
}
