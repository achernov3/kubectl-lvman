package commands

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"
	"kubectl-lvman/internal/k8s"
	"kubectl-lvman/internal/table"
	"slices"

	"github.com/urfave/cli/v2"
)

var (
	orphan = &cli.Command{
		Name:    config.CmdOrphan,
		Aliases: []string{"o"},
		Usage:   "prints oprhaned logical volumes and the nodes on which they are located",
		Flags:   config.OrphanFlags,
		Action:  showOrphan,
	}
)

func showOrphan(clictx *cli.Context) error {
	var pvs []string
	var tableData [][]string
	tableRender := table.GetTableRender()
	ctx := context.Background()

	cfg, err := config.NewConfig(clictx)
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
			tableData = append(tableData, []string{lvName, lv.Spec.NodeName, lv.Status.VolumeID})
		}

	}

	tableRender.RenderTable(tableData, []string{"LogicalVolume", "NODE", "VOLUME ID"})

	return nil

}
