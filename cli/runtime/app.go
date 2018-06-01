package runtime

import (
	"errors"
	"fmt"
	"log"

	"github.com/urfave/cli"
)

const (
	maxNumberOfArguments = 9
)

var defaultKeyVal = make(map[string]string, maxNumberOfArguments)

var defaultAppOptions = options{
	osArgs:    []string{},
	arguments: defaultKeyVal,
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type App struct {
	opts options
}

func (runtime *App) Run(app *cli.App) (err error) {

	err = app.Run(runtime.opts.osArgs)
	if err != nil {
		log.Fatal(fmt.Sprintf("CLI Run command failed: %v", err))
	}

	return
}

type options struct {
	osArgs    []string
	arguments map[string]string
	testFn    testFunc
}

func New(appOpts ...AppOption) (*App, error) {
	opts := defaultAppOptions
	for _, f := range appOpts {
		err := f(&opts)
		if err != nil {
			//return nil, errors.Wrap(err, "error setting option")
			return nil, errors.New("error setting option")
		}
	}

	// Copy arguments to unusedArguments
	unusedArguments := make(map[string]string)
	for k, v := range opts.arguments {
		unusedArguments[k] = v
	}

	for i, arg := range opts.osArgs {
		// if arg found in opts.arguments replace arg value
		if opts.arguments[arg] != "" {
			if len(opts.osArgs) >= i {
				opts.osArgs[i+1] = opts.arguments[arg]
				delete(unusedArguments, arg)
			} else {
				fmt.Printf("len(opts.osArgs) >= i is false. Unable to replace %s\n", opts.arguments[arg])
			}
		}
	}

	if len(unusedArguments) > 0 {
		// Add executable name
		updatedArgs := []string{opts.osArgs[0]}

		// Add unused arguments: These are new arguments that were added programmatically
		for k, v := range unusedArguments {
			updatedArgs = append(updatedArgs, k)
			updatedArgs = append(updatedArgs, v)
		}

		// Add all arguments from original os args slice
		updatedArgs = append(updatedArgs, opts.osArgs[1:]...)

		// Replace opts with new list of arguments containing new flags
		opts.osArgs = updatedArgs
	}

	s := &App{
		opts: opts,
	}
	return s, nil
}
