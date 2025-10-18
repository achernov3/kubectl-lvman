package commands

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"
	diskfree "kubectl-lvman/internal/disk_free"
	"kubectl-lvman/internal/k8s"
	sshclient "kubectl-lvman/internal/ssh_client"
	"kubectl-lvman/internal/table"

	"github.com/urfave/cli/v3"
	v1 "k8s.io/api/core/v1"
)

var (
	diskFree = &cli.Command{
		Name:  config.CmdDF,
		Usage: "prints disk usage to stdout and other info about pvc, pv, lv",
		Flags: config.ShowFlags,
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			if c.Name == config.CmdDF && len(c.Args().Slice()) == 0 {
				return nil, fmt.Errorf("you must specify pvc names!")
			}
			return nil, nil
		},
		Action: showDiskFree,
	}

	standardHeader []string = []string{"PVC", "PV", "STATUS", "NODE", "VOLUME ID", "CAPACITY", "USAGE"}
)

func showDiskFree(ctx context.Context, cmd *cli.Command) error {
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

	for _, pvcName := range cfg.PVCNames {
		pvc, err := client.GetPVC(cfg.Namespace, pvcName, ctx)
		if err != nil {
			return fmt.Errorf("can't get resource *PersistentVolumeClaim: %w", err)
		}

		pv, err := client.GetPV(pvc.Spec.VolumeName, ctx)
		if err != nil {
			return fmt.Errorf("can't get resource *PersistentVolume: %w", err)
		}

		if pv.Spec.NodeAffinity == nil ||
			pv.Spec.NodeAffinity.Required == nil ||
			len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms) == 0 ||
			len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions) == 0 {
			return fmt.Errorf("PV %s has no node affinity configuration", pv.Name)
		}

		node, err := client.GetNode(pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values[0], ctx)
		if err != nil {
			return fmt.Errorf("can't get resource *Node: %w", err)
		}

		var host string
		for _, node := range node.Status.Addresses {
			if node.Type == v1.NodeInternalIP {
				host = node.Address
			}
		}

		cmd := fmt.Sprintf(`df -hP | grep "%s" | awk '
			BEGIN {
			    printf "{\"discarray\":["
			}
			{
			    if ($1 == "Filesystem") next
			    if (a) printf ","
			    printf "{\"mount\":\"" $6 "\",\"size\":\"" $2 "\",\"used\":\"" $3 "\",\"avail\":\"" $4 "\",\"use_percent\":\"" $5 "\"}"
			    a++
			}
			END {
			    print "]}"
			}'`, pv.Spec.CSI.VolumeHandle)

		stdout, err := sshclient.ExecCMD(cfg, cmd, host)
		if err != nil {
			return fmt.Errorf("df command exit with non zero: %w", err)
		}

		df, err := diskfree.UnmarshalJSON(stdout)
		if err != nil {
			return fmt.Errorf("can't unmarshal stdout: %w", err)
		}

		tableData = append(tableData, table.MakeColumnsSlice(pvc, pv, node, *df))
	}

	tableRender.RenderTable(tableData, standardHeader)

	return nil
}
