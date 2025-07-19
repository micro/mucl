package main

import (
	"github.com/micro/mucl/cmd"

	// load packages so they can register commands
	_ "github.com/micro/mucl/cmd/ebnf"
	_ "github.com/micro/mucl/cmd/gen"
	_ "github.com/micro/mucl/cmd/init"
)

func main() {
	cmd.Run()
}
