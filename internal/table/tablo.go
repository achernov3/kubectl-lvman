package table

import (
	diskfree "kubectl-lvman/internal/disk_free"

	v1 "k8s.io/api/core/v1"
)

func MakeColumnsSlice(pvc *v1.PersistentVolumeClaim, pv *v1.PersistentVolume, node *v1.Node, d diskfree.DiskFree) []string {
	var usage string

	if len(d.Discarray) == 0 {
		usage = "UNMOUNT"
	} else {
		usage = d.Discarray[0].Used
	}

	return []string{pvc.Name, pv.Name,
		string(pv.Status.Phase),
		node.Name,
		pv.Spec.CSI.VolumeHandle,
		pv.Spec.Capacity.Storage().String(),
		usage}
}
