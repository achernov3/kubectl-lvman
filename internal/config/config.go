package config

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"k8s.io/client-go/util/homedir"
)

const (
	argKubeconfig = "kubeconfig"
	argNamespace  = "namespace"
	argContext    = "context"
	argIdRsa      = "id_rsa"
	argUsername   = "username"
	argPort       = "port"
	ArgFormat     = "format"
)

const (
	CmdDF          = "df"
	CmdShow        = "show"
	CmdShowShort   = "s"
	CmdOrphan      = "orphan"
	CmdOrphanShort = "o"
	CmdPrune       = "prune"
	CmdPruneShort  = "p"
	CmdLost        = "lost"
	CmdLostShort   = "l"
)

var (
	ShowOrphanHeaders = []string{"LOGICAL VOLUME", "NODE", "VOLUME ID"}
	StandardHeader    = []string{"PVC", "PV", "STATUS", "NODE", "VOLUME ID", "CAPACITY", "USAGE"}
)

var (
	ShowFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        argKubeconfig,
			Usage:       "kubernetes client config path",
			Sources:     cli.EnvVars("KUBECONFIG"),
			Value:       fmt.Sprintf("%s/.kube/config", homedir.HomeDir()),
			Validator:   validateStringFlagsNonEmpty,
			DefaultText: "$HOME/.kube/config",
		},
		&cli.StringFlag{
			Name:      argNamespace,
			Aliases:   []string{"n"},
			Usage:     "override namespace of current context from kubeconfig",
			Value:     "",
			Validator: validateStringFlagsNonEmpty,
		},
		&cli.StringFlag{
			Name:      argContext,
			Usage:     "override current context from kubeconfig",
			Value:     "",
			Validator: validateStringFlagsNonEmpty,
		},
		&cli.StringFlag{
			Name:        argIdRsa,
			Usage:       "Path to private ssh key",
			Value:       fmt.Sprintf("%s/.ssh/id_rsa", homedir.HomeDir()),
			Validator:   validateStringFlagsNonEmpty,
			DefaultText: "$HOME/.ssh/id_rsa",
		},
		&cli.StringFlag{
			Name:      argUsername,
			Usage:     "Paste username for ssh to node",
			Value:     "ops",
			Validator: validateStringFlagsNonEmpty,
		},
		&cli.StringFlag{
			Name:      argPort,
			Usage:     "Paste port for ssh connection",
			Value:     "22",
			Validator: validateStringFlagsNonEmpty,
		},
	}

	OrphanFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        argKubeconfig,
			Usage:       "kubernetes client config path",
			Sources:     cli.EnvVars("KUBECONFIG"),
			Value:       fmt.Sprintf("%s/.kube/config", homedir.HomeDir()),
			Validator:   validateStringFlagsNonEmpty,
			DefaultText: "$HOME/.kube/config",
		},
		&cli.StringFlag{
			Name:      argNamespace,
			Aliases:   []string{"n"},
			Usage:     "override namespace of current context from kubeconfig",
			Value:     "",
			Validator: validateStringFlagsNonEmpty,
		},
		&cli.StringFlag{
			Name:      argContext,
			Usage:     "override current context from kubeconfig",
			Value:     "",
			Validator: validateStringFlagsNonEmpty,
		},
		&cli.BoolFlag{
			Name:    ArgFormat,
			Aliases: []string{"r"},
			Usage:   "allows show stdout in raw format",
		},
	}
)

type Config struct {
	KubeConfig  string
	Namespace   string
	KubeContext string
	PVCNames    []string
	SSHKey      string
	Username    string
	Port        string
}

func NewConfig(ctx context.Context, cmd *cli.Command) (*Config, error) {
	return &Config{
		KubeConfig:  cmd.String(argKubeconfig),
		Namespace:   cmd.String(argNamespace),
		KubeContext: cmd.String(argContext),
		PVCNames:    cmd.Args().Slice(),
		SSHKey:      cmd.String(argIdRsa),
		Username:    cmd.String(argUsername),
		Port:        cmd.String(argPort),
	}, nil
}

func validateStringFlagsNonEmpty(s string) error {
	if s == "" {
		return fmt.Errorf("option --%s must not be empty", s)
	}
	return nil
}
