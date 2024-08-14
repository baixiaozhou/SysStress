package cmd

import "github.com/urfave/cli/v2"

const (
	version = "1.0.0"
)

func Main(args []string) error {
	app := &cli.App{
		Name:            "sysstress",
		Usage:           "A powerful tool designed to simulate high loads on your system",
		Version:         version,
		HideHelpCommand: true,
		Commands: []*cli.Command{
			stressCpu(),
		},
	}
	return app.Run(args)
}
