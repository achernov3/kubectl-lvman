package commands

import (
	"kubectl-lvman/internal/config"
	"kubectl-lvman/internal/subcmd"

	"github.com/urfave/cli/v3"
)

var (
	Show = &cli.Command{
		Name:    config.CmdShow,
		Aliases: []string{config.CmdShowShort},
		Usage:   "shows disk usage by provided pvc names and oprhaned logical volumes in cluster",
		Commands: []*cli.Command{
			subcmd.DF,
			subcmd.Orphan,
		},
	}
)
