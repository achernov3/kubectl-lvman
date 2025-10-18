package commands

import (
	"kubectl-lvman/internal/config"

	"github.com/urfave/cli/v3"
)

var (
	Show = &cli.Command{
		Name:    config.CmdShow,
		Aliases: []string{config.CmdShowShort},
		Usage:   "",
		Commands: []*cli.Command{
			diskFree,
			orphan,
		},
	}
)
