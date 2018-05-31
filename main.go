package main

import (
	"log"
	"os"

	"github.com/l3x/mycli/myDir/cli"
)

func main() {
	// Grab the user inputed CLI flags
	cliFlags := cli.FlagsStruct{}
	cliErr := cli.StartCLI(&cliFlags, os.Args)
	if cliErr != nil {
		log.Println("Error grabbing command line args")
		log.Fatal(cliErr)
	}

	// Do stuff ...
}