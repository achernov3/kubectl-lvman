package main

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/commands"
	"kubectl-lvman/internal/config"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

var Version = "local"

func main() {
	app := &cli.Command{
		Name:                   config.AppName,
		Version:                Version,
		Metadata:               map[string]interface{}{"Compiled:": time.Now()},
		Usage:                  "kubectl plugin for managing logical volumes in a kubernetes cluster with topolvm as storage class",
		UsageText:              fmt.Sprintf(`%s [flags] [command]`, config.AppName),
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		HideHelpCommand:        true,
		Commands: []*cli.Command{
			commands.Show,
			commands.Prune,
			commands.Remove,
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("%+v: %+v", os.Args[0], err)
	}
}
