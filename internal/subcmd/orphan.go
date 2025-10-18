package subcmd

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"
	"kubectl-lvman/internal/k8s"
	"kubectl-lvman/internal/table"
	"slices"

	"github.com/urfave/cli/v3"
)

var (
	Orphan = &cli.Command{
		Name:    config.CmdOrphan,
		Aliases: []string{config.CmdOrphanShort},
		Usage:   "prints oprhaned logical volumes and the nodes on which they are located",
		Flags: []cli.Flag{
			config.KubeConfigFlag[0],
			config.KubeContextFlag[0],
			config.KubeNamespaceFlag[0],
		},
		Action:    showOrphan,
		UsageText: fmt.Sprintf(`%s %s %s [flags] [command]`, config.AppName, config.CmdShow, config.CmdOrphan),
	}
)

func showOrphan(ctx context.Context, cmd *cli.Command) error {
	var pvs []string
	var tableData [][]string
	tableRender := table.GetTableRender()

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
			tableData = append(tableData, []string{lvName, lv.Spec.NodeName, lv.Status.VolumeID})
		}
	}

	if len(tableData) == 0 {
		fmt.Println("There's no oprhaned logical volumes")
	} else {
		tableRender.RenderTable(tableData, config.ShowOrphanHeaders)
	}

	return nil

}
