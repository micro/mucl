package main

import (
	"github.com/micro/mu/cmd"

	// load packages so they can register commands
	_ "github.com/micro/mu/cmd/gen"
)

func main() {
	cmd.Run()
}
