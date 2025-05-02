package gen

import (
	"fmt"
	"os"

	"strings"

	"github.com/charmbracelet/huh"
	"github.com/micro/mu/cmd"
	"github.com/micro/mu/generator"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "init",
			Usage:  "create an initial mu configuration",
			Action: Run,
			Flags:  Flags,
		},
	)
}

func Run(c *cli.Context) error {
	var serviceName string
	var endpointName string
	var rpcMethodName string
	var goModuleName string
	err := huh.NewForm(
		huh.NewGroup(

			huh.NewInput().Title("Service Name").
				Placeholder("Name of the service. e.g. users").
				Value(&serviceName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("service name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("service name cannot contain spaces")
					}
					return nil
				}),
			huh.NewInput().Title("Endpoint name").
				Placeholder("Name of the Endpoint. e.g. User").
				Value(&endpointName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("endpoint name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("endpoint name cannot contain spaces")
					}
					return nil
				}),
			huh.NewInput().Title("Method name").
				Placeholder("Name of the first method. e.g. Create").
				Value(&rpcMethodName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("method name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("method name cannot contain spaces")
					}
					return nil
				}),
			huh.NewInput().Title("Go Module name").
				Placeholder("Name of the first method. e.g. userservice or github.com/myorg/userservice").
				Value(&goModuleName).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("module name cannot be empty")
					}
					if strings.Contains(s, " ") {
						return fmt.Errorf("module name cannot contain spaces")
					}
					return nil
				}),
		).Title("Create a new go-micro service definition"),
	).
		WithTheme(huh.ThemeBase()).Run()
	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		fmt.Println(err)
		os.Exit(1)
	}

	// Create the mu file
	return generator.CreateConfig(serviceName, endpointName, rpcMethodName, goModuleName, c.String("definition"))

}

var (
	Flags = []cli.Flag{

		&cli.StringFlag{
			Name:    "definition",
			Usage:   "mu definition file",
			EnvVars: []string{"MU_DEFINITION"},
			Value:   "service.mu",
		},
	}
)
