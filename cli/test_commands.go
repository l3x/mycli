package cli

import (
	"errors"

	"os"

	"fmt"

	"github.com/urfave/cli"
)

func TestCommands(c *cli.Context) (err error) {

	testName := os.Getenv("TEST_NAME")
	if testName == "test-extra-arg" {
		err = testExtraArg(c)
	}

	return
}

func testExtraArg(c *cli.Context) (err error) {
	want := os.Getenv("TEST_WANT")
	got := c.GlobalString("extra")

	if want != got {
		err = errors.New(fmt.Sprintf("expected extra flag to equal (%s), but got (%s)", want, got))
	}

	return
}
