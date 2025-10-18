package main

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/commands"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

var Version = "local"

const appName = "kubectl-lvman"

func main() {
	app := &cli.Command{
		Name:                   appName,
		Version:                Version,
		Metadata:               map[string]interface{}{"Compiled:": time.Now()},
		Usage:                  "kubectl plugin for managing logical volumes in a kubernetes cluster with topolvm as storage class",
		UsageText:              fmt.Sprintf(`%s [flags] [command]`, appName),
		UseShortOptionHandling: true,
		EnableShellCompletion:  true,
		HideHelpCommand:        true,
		Commands: []*cli.Command{
			commands.Show,
			commands.Prune,
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("%+v: %+v", os.Args[0], err)
	}
}
