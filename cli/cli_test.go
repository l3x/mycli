package cli

import (
	"os"
	"testing"

	"github.com/l3x/mycli/cli/runtime"
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
		testName string
		args     map[string]string
		want     string
	}

	data := []testData{
		{
			"test-extra-arg",
			map[string]string{
				"--config": "/tmp/xxx",
				"--extra":  "XXX",
			},
			"XXX",
		},
		{
			"test-extra-arg",
			map[string]string{
				"--extra": "zzz",
			},
			"zzz",
		},
	}

	for _, test := range data {

		os.Setenv("TEST_NAME", test.testName)
		os.Setenv("TEST_WANT", test.want)
		RunCLI(osArgs, Version, Revision, runtime.Argument(test.args))

	}

}
