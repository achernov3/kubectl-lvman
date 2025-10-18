package commands

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"
	"kubectl-lvman/internal/k8s"
	"slices"

	"github.com/urfave/cli/v3"
)

var (
	Prune = &cli.Command{
		Name:    config.CmdPrune,
		Aliases: []string{config.CmdPruneShort},
		Usage:   "prune all oprhaned LV (which hasn't binded PV)",
		Action:  pruneOrphan,
		Flags: []cli.Flag{
			config.KubeConfigFlag[0],
			config.KubeContextFlag[0],
			config.KubeNamespaceFlag[0],
		},
	}
)

func pruneOrphan(ctx context.Context, cmd *cli.Command) error {
	var pvs []string

	cfg, err := config.NewConfig(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't get *Config: %w", err)
	}

	client, err := k8s.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("can't get *Client: %w", err)
	}

	lvList, err := client.ListLV(ctx)
	if err != nil {
		return fmt.Errorf("can't get resource LogicalVolumeList: %w", err)
	}

	pvList, err := client.ListPV(ctx)
	if err != nil {
		return fmt.Errorf("can't get resource PersistentVolumeList: %w", err)
	}

	for _, pv := range pvList.Items {
		pvs = append(pvs, pv.GetName())
	}

	for _, lv := range lvList.Items {
		lvName := lv.GetName()

		if !slices.Contains(pvs, lvName) {
			err = client.DeleteLV(lvName, ctx)
			if err != nil {
				return fmt.Errorf("failed to delete LogicalVolume %v: %w", lvName, err)
			}
		}
	}

	return nil
}
