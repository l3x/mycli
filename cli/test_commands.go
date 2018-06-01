package cli

import (
	"errors"

	"github.com/urfave/cli"
)

func TestCommands(c *cli.Context) (err error) {

	err = testExtraForXXX(c)

	return
}

func testExtraForXXX(c *cli.Context) (err error) {
	want := "xxx"
	extra := c.GlobalString("extra")

	if extra != want {
		err = errors.New("expected extra flag to equal " + want)
	}

	return
}
