package cli

import "github.com/urfave/cli"

// FlagsStruct - holds command line args
type FlagsStruct struct {
	MyFlag string
}

// StartCLI - gathers command line args
func StartCLI(cliFlags *FlagsStruct, args []string) error {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		MyFlag := ctx.GlobalString("my-flag")

		// build the cli struct to send back to main
		cliFlags.MyFlag = MyFlag
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
			Name:  "my-flag, f",
			Usage: "My flag usage goes here",
			Value: "myDefault",
		},
	}
	app.Name = "myAppName"
	app.Usage = "My App's Usage"
	app.Version = "0.0.1"
	return app.Run(args)
}