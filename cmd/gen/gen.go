package gen

import (
	"errors"
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
	f := c.Args().First()
	if f == "" {
		return errors.New("please provide an input file as the first argument")
	}
	// Read the file

	bb, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	def, err := mucl.Parser.ParseBytes(f, bb)
	if err != nil {
		return fmt.Errorf("parsing failure: %v", err)
	}
	g, err := generator.NewGenerator(def)
	if err != nil {
		return fmt.Errorf("failed to create generator: %v", err)
	}
	if err := g.Generate(); err != nil {
		return fmt.Errorf("failed to generate: %v", err)
	}
	return nil
}

var (
	Flags = []cli.Flag{}
)
