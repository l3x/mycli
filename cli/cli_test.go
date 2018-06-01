package cli

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/l3x/mycli/cli/runtime"
	"github.com/urfave/cli"
)

const (
	defaultVerboseLogging = false
)

var (
	Version        = "1.0.0"
	Revision       = "591fbe1"
	cliBinaryName  = "keep-core"
	osArgs         = []string{cliBinaryName, "test-commands"}
	verboseLogging bool
)

//-------------------------------------------------------------------------------
// Main
//-------------------------------------------------------------------------------
func TestMain(m *testing.M) {
	setup()

	code := m.Run()
	os.Exit(code)

	teardown()
}

func setup() {
	////TODO: Remove after testing on circleci
	//cmd := exec.Command("sh", "-c", "find ../../.. -name \"keep-*\"")
	//stdoutStderr, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("Error running command: %v\n", err)
	//} else {
	//	fmt.Printf("%s\n", stdoutStderr)
	//}

	//cliBinaryName = os.Getenv("CLI_BINARY")
	//if len(cliBinaryName) == 0 {
	//	cliBinaryName = defaultCLIBinaryName
	//}
	//if !FileExists(cliBinaryName) {
	//	fmt.Printf("CLI binary file not found: %s\n", cliBinaryName)
	//	os.Exit(1)
	//}

	verboseLogging = defaultVerboseLogging
	if os.Getenv("CLI_VERBOSE_LOGGING") == "true" {
		verboseLogging = true
	}
}

func teardown() {
	// global test teardown instructions
}

//-------------------------------------------------------------------------------
// Tests
//-------------------------------------------------------------------------------

func TestConfigFlag(t *testing.T) {

	type testData struct {
		args map[string]string
		want string
	}

	data := []testData{
		{map[string]string{
			"--config": "/tmp/xxx",
			"--extra":  "XXX",
		}, "XXX"},
	}

	for _, test := range data {
		newArgs := runtime.Argument(test.args)

		runtime.TestFunc(func(c *cli.Context) bool { return c.GlobalString("extra") == "XXX" })

		cliErr := RunCLI(osArgs, Version, Revision, newArgs)
		if cliErr != nil {
			log.Println("CLI error encountered:")
			log.Fatal(cliErr)
		}

		set := flag.NewFlagSet("testExtraForXXX", flag.ContinueOnError)
		app := cli.NewApp()
		ctx := cli.NewContext(app, set, nil)

		got := ctx.GlobalString("extra")

		if test.want != got {
			t.Errorf("%s %s => \nWANT\n%s\nGOT\n%v", test.want, got)
		}
	}

}
