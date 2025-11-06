package commands

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"
	"kubectl-lvman/internal/k8s"

	"github.com/urfave/cli/v3"
)

var (
	Remove = &cli.Command{
		Name:      config.CmdRemove,
		Aliases:   []string{config.CmdRemoveShort},
		Usage:     "remove chain pvc -> pv -> lv by provided pvc",
		UsageText: fmt.Sprintf(`%s %s [flags] [command] <pvc-list>`, config.AppName, config.CmdRemove),
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			if c.Name == config.CmdRemove && len(c.Args().Slice()) == 0 {
				return nil, fmt.Errorf("you must specify pvc names")
			}
			return nil, nil
		},
		Action: removeChain,
		Flags: []cli.Flag{
			config.KubeConfigFlag[0],
			config.KubeContextFlag[0],
			config.KubeNamespaceFlag[0],
		},
	}
)

func removeChain(ctx context.Context, cmd *cli.Command) error {
	cfg, err := config.NewConfig(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't get *Config: %w", err)
	}

	client, err := k8s.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("can't get *Client: %w", err)
	}

	for _, pvcName := range cfg.PVCNames {
		pvc, err := client.GetPVC(cfg.Namespace, pvcName, ctx)
		if err != nil {
			return fmt.Errorf("can't get resource *PersistentVolumeClaim: %w", err)
		}

		err = client.DeletePV(pvc.Spec.VolumeName, ctx)
		if err != nil {
			return fmt.Errorf("can't delete resource *PersistentVolume %s: %w", pvc.Spec.VolumeName, err)
		}

		err = client.DeleteLV(pvc.Spec.VolumeName, ctx)
		if err != nil {
			return fmt.Errorf("can't delete resource *LogicalVolume %s: %w", pvc.Spec.VolumeName, err)
		}
	}

	return nil
}
