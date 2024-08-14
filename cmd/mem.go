package cmd

import "github.com/urfave/cli/v2"

func stressMem() *cli.Command {
	return &cli.Command{
		Name:  "memory",
		Usage: "memory stress",
		Description: `This command stress memory stress.
`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "size",
				Aliases:  []string{"s"},
				Usage:    "memory size",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "duration",
				Aliases:  []string{"d"},
				Usage:    "memory duration",
				Required: true,
			},
		},
		Action: memory,
	}
}

func memory(c *cli.Context) error {
	return nil
}
