package gen

import (
	"fmt"
	"os"

	"github.com/micro/mu/cmd"
	"github.com/micro/mu/generator"
	"github.com/micro/mu/mucl"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "gen",
			Usage:  "generate code from a mucl file",
			Action: Run,
			Flags:  Flags,
		},
	)
}

func Run(c *cli.Context) error {
	f := c.String("definition")
	if f == "" {
		return fmt.Errorf("definition file not provided")
	}
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return fmt.Errorf("definition file does not exist: %s", f)
	}
	bb, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	def, err := mucl.Parser.ParseBytes(f, bb)
	if err != nil {
		return fmt.Errorf("parsing failure: %v", err)
	}

	onlyTypes := c.Bool("types")

	g, err := generator.NewGenerator(def, onlyTypes)
	if err != nil {
		return fmt.Errorf("failed to create generator: %v", err)
	}
	if err := g.Generate(); err != nil {
		return fmt.Errorf("failed to generate: %v", err)
	}
	return nil
}

var (
	Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "types",
			Usage:   "only generate types",
			EnvVars: []string{"MU_GEN_TYPES"},
		},
		&cli.StringFlag{
			Name:    "definition",
			Usage:   "mu definition file",
			EnvVars: []string{"MU_DEFINITION"},
			Value:   "service.mu",
		},
	}
)
