package command

import (
	"github.com/urfave/cli/v2"
	containercommand "raind/internal/command/container"
	imagecommand "raind/internal/command/image"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "raind",
		Usage: "raind container runtime",
		Commands: []*cli.Command{
			{
				Name:  "container",
				Usage: "container operation",
				Subcommands: []*cli.Command{
					containercommand.CommandCreate(),
					containercommand.CommandStart(),
					containercommand.CommandStop(),
					containercommand.CommandRemove(),
					containercommand.CommandList(),
					containercommand.CommandAttach(),
					containercommand.CommandRun(),
				},
			},
			{
				Name:  "image",
				Usage: "image operation",
				Subcommands: []*cli.Command{
					imagecommand.CommandPull(),
					imagecommand.CommandList(),
					imagecommand.CommandRemove(),
				},
			},
		},
	}

	// disable slice flag separator
	app.DisableSliceFlagSeparator = true

	return app
}
