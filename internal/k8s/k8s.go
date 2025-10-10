package k8s

import (
	"context"
	"fmt"
	"kubectl-lvman/internal/config"

	// старое деприкейтнутое говно

	topolvmlegacyv1 "github.com/topolvm/topolvm/api/legacy/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const logicalvolumes = "logicalvolumes"

func (c *Client) GetPVC(namespace, pvcName string, ctx context.Context) (*v1.PersistentVolumeClaim, error) {
	return c.Clientset.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, pvcName, metav1.GetOptions{})
}

func (c *Client) GetPV(volumeName string, ctx context.Context) (*v1.PersistentVolume, error) {
	return c.Clientset.CoreV1().PersistentVolumes().Get(ctx, volumeName, metav1.GetOptions{})

}

func (c *Client) GetNode(nodeName string, ctx context.Context) (*v1.Node, error) {
	return c.Clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
}

func (c *Client) ListPV(ctx context.Context) (*v1.PersistentVolumeList, error) {
	return c.Clientset.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
}

func (c *Client) GetLV(cfg *config.Config, name string, ctx context.Context) (*topolvmlegacyv1.LogicalVolume, error) {
	resourceScheme := topolvmlegacyv1.SchemeBuilder.GroupVersion.WithResource(logicalvolumes)

	LogicalVolumeResource := &topolvmlegacyv1.LogicalVolume{}

	u, err := c.DynamicClient.Resource(resourceScheme).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't get Unstructured: %w", err)
	}

	unstructured := u.UnstructuredContent()

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, LogicalVolumeResource)
	if err != nil {
		return nil, fmt.Errorf("failed to convert unstructured to LogicalVolumeResource: %w", err)
	}

	return LogicalVolumeResource, err
}

func (c *Client) ListLV(ctx context.Context) (*topolvmlegacyv1.LogicalVolumeList, error) {
	resourceScheme := topolvmlegacyv1.SchemeBuilder.GroupVersion.WithResource(logicalvolumes)

	LogicalVolumeListResource := &topolvmlegacyv1.LogicalVolumeList{}

	u, err := c.DynamicClient.Resource(resourceScheme).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("can't get UnstructuredList: %w", err)
	}

	unstructured := u.UnstructuredContent()

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, LogicalVolumeListResource)
	if err != nil {
		return nil, fmt.Errorf("failed to convert unstructured to LogicalVolumeListResource: %w", err)
	}

	return LogicalVolumeListResource, err
}
