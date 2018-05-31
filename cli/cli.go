package cli

import "github.com/urfave/cli"

// FlagsStruct - holds command line args
type FlagsStruct struct {
	ConfigFlag string
}

// StartCLI - gathers command line args
func StartCLI(cliFlags *FlagsStruct, args []string) error {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		ConfigFlag := ctx.GlobalString("config")

		// build the cli struct to send back to main
		cliFlags.ConfigFlag = ConfigFlag
		return nil
	}
	app.Authors = []cli.Author{
		{
			Email: "my@email.com",
			Name:  "Adam Hanna",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Config flag usage goes here",
			Value: "../test/config.toml",
		},
	}
	app.Name = "myAppName"
	app.Usage = "My App's Usage"
	app.Version = "0.0.1"
	return app.Run(args)
}