package commands

import (
	"kubectl-lvman/internal/config"

	"github.com/urfave/cli/v2"
)

var (
	Show = &cli.Command{
		Name:    config.CmdShow,
		Aliases: []string{"s"},
		Usage:   "",
		Subcommands: []*cli.Command{
			diskFree,
			orphan,
		},
	}
)
