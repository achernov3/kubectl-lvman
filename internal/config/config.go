package config

import (
	"fmt"

	"github.com/urfave/cli/v2"
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
	CmdDF     = "df"
	CmdShow   = "show"
	CmdOrphan = "orphan"
	CmdPrune  = "prune"
)

var (
	ShowFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        argKubeconfig,
			Usage:       "kubernetes client config path",
			EnvVars:     []string{"KUBECONFIG"},
			Value:       fmt.Sprintf("%s/.kube/config", homedir.HomeDir()),
			DefaultText: "$HOME/.kube/config",
		},
		&cli.StringFlag{
			Name:    argNamespace,
			Aliases: []string{"n"},
			Usage:   "override namespace of current context from kubeconfig",
			Value:   "",
		},
		&cli.StringFlag{
			Name:  argContext,
			Usage: "override current context from kubeconfig",
			Value: "",
		},
		&cli.PathFlag{
			Name:        argIdRsa,
			Usage:       "Path to private ssh key",
			Value:       fmt.Sprintf("%s/.ssh/id_rsa", homedir.HomeDir()),
			DefaultText: "$HOME/.ssh/id_rsa",
		},
		&cli.StringFlag{
			Name:  argUsername,
			Usage: "Paste username for ssh to node",
			Value: "ops",
		},
		&cli.StringFlag{
			Name:  argPort,
			Usage: "Paste port for ssh connection",
			Value: "22",
		},
	}

	OrphanFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        argKubeconfig,
			Usage:       "kubernetes client config path",
			EnvVars:     []string{"KUBECONFIG"},
			Value:       fmt.Sprintf("%s/.kube/config", homedir.HomeDir()),
			DefaultText: "$HOME/.kube/config",
		},
		&cli.StringFlag{
			Name:    argNamespace,
			Aliases: []string{"n"},
			Usage:   "override namespace of current context from kubeconfig",
			Value:   "",
		},
		&cli.StringFlag{
			Name:  argContext,
			Usage: "override current context from kubeconfig",
			Value: "",
		},
		&cli.BoolFlag{
			Name:    ArgFormat,
			Aliases: []string{"r"},
			Usage:   "allows show stdout in raw format",
		},
	}

	stringFlags = []string{argKubeconfig, argNamespace, argContext}
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

func NewConfig(clictx *cli.Context) (*Config, error) {
	pvcName := clictx.Args().Slice()
	if clictx.Command.Name != CmdOrphan {
		if len(pvcName) == 0 {
			return nil, errorWithCliHelp(clictx, "you must specify pvc name!")
		}
	}

	err := validateStringFlagsNonEmpty(clictx, stringFlags)
	if err != nil {
		return nil, err
	}

	return &Config{
		KubeConfig:  clictx.String(argKubeconfig),
		Namespace:   clictx.String(argNamespace),
		KubeContext: clictx.String(argContext),
		PVCNames:    pvcName,
		SSHKey:      clictx.Path(argIdRsa),
		Username:    clictx.String(argUsername),
		Port:        clictx.String(argPort),
	}, nil
}

func errorWithCliHelp(clictx *cli.Context, a any) error {
	err := cli.ShowAppHelp(clictx)
	if err != nil {
		return err
	}
	//nolint:staticcheck
	return fmt.Errorf("%s\n", a)
}

func errorWithCliHelpf(clictx *cli.Context, format string, a ...any) error {
	return errorWithCliHelp(clictx, fmt.Sprintf(format, a...))
}

func validateStringFlagsNonEmpty(clictx *cli.Context, flags []string) error {
	for _, flag := range flags {
		if clictx.IsSet(flag) {
			if clictx.String(flag) == "" {
				return errorWithCliHelpf(clictx, "option --%s must not be empty", flag)
			}
		}
	}
	return nil
}
