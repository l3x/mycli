package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/urfave/cli"
)

const (
	defaultVerboseLogging = false
	testFileMode          = os.FileMode(0640)
)

var (
	cliBinaryName  string
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

	var Version = "1.0.0"
	var Revision = "591fbe1"

	arguments := []string{
		"/private/var/folders/8g/tm3k26dx3x30w7x0lhq_2qg00000gn/T/___go_build_main_go_darwin",
		"--config",
		"/tmp/UTC--2018-03-11T01-37-33.202765887Z--c2a56884538778bacd91aa5bf343bf882c5fb18b",
		"test-commands",
	}
	cliErr := RunCLI(arguments, Version, Revision)
	if cliErr != nil {
		log.Println("CLI error encountered:")
		log.Fatal(cliErr)
	}

	set := flag.NewFlagSet("testFlags", flag.ContinueOnError)
	err := set.Parse(arguments[1:])
	if err != nil {
		t.Errorf("error parsing arguments: %v", err)
	}
	app := cli.NewApp()
	ctx := cli.NewContext(app, set, nil)

	fmt.Printf("ctx.App.Flags: %v\n", ctx.App.Flags)

	//os.Setenv("KEEP_ETHEREUM_PASSWORD", test.config.Ethereum.Account.KeyFilePassword)

	//err := validateConfig(test.config)
	//var got string
	//if err != nil {
	//	got = strings.TrimSpace(err.Error())
	//}

	//if test.want != got {
	//	t.Errorf("%s %s => \nWANT\n%s\nGOT\n%v", test.want, got)
	//}

}
