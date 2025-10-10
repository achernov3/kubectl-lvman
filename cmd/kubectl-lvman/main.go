package main

import (
	"fmt"
	"kubectl-lvman/internal/commands"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var Version = "local"

const appName = "kubectl-lvman"

func main() {
	app := &cli.App{
		Name:     appName,
		Version:  Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "anon",
			},
		},
		HelpName: appName,
		Usage:    "kubectl plugin for managing logical volumes in a cluster",
		UsageText: fmt.Sprintf(`%s [flags] [command]

		Example:
		
		%s show --username admin --port 2221 --id_rsa ~/.ssh/admin/id_rsa show df <pvc-names>

		%s show orphan
		`, appName, appName, appName),
		UseShortOptionHandling: true,
		EnableBashCompletion:   true,
		HideHelpCommand:        true,
		Commands: []*cli.Command{
			commands.Show,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%+v: %+v", os.Args[0], err)
	}
}
