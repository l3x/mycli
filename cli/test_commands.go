package cli

import (
	"log"

	"fmt"

	"github.com/urfave/cli"
)

func TestCommands(c *cli.Context) (err error) {

	testFlags(c)

	return
}

func testFlags(c *cli.Context) {
	arguments := []string{
		"/private/var/folders/8g/tm3k26dx3x30w7x0lhq_2qg00000gn/T/___go_build_main_go_darwin",
		"--config",
		"/tmp/UTC--2018-03-11T01-37-33.202765887Z--c2a56884538778bacd91aa5bf343bf882c5fb18b",
		"test-flags",
	}
	var Version = "1.0.0"
	var Revision = "591fbe1"
	cliErr := RunCLI(arguments, Version, Revision)
	if cliErr != nil {
		log.Println("CLI error encountered:")
		log.Fatal(cliErr)
	}
	configPath := c.GlobalString("config")
	fmt.Printf("configPath: %v\n", configPath)

}
