package k8s

import (
	"fmt"
	"kubectl-lvman/internal/config"
	"strings"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	Clientset     *kubernetes.Clientset
	DynamicClient *dynamic.DynamicClient
}

func NewClient(cfg *config.Config) (*Client, error) {
	Clientset, err := getKubernetesClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("can't build client: %w", err)
	}

	DynamicClient, err := getDynamicClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("can't build dynamic client: %w", err)
	}

	return &Client{
		Clientset:     Clientset,
		DynamicClient: DynamicClient,
	}, nil
}

func getKubernetesClient(cfg *config.Config) (*kubernetes.Clientset, error) {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{Precedence: strings.Split(cfg.KubeConfig, ":")},
		&clientcmd.ConfigOverrides{
			CurrentContext: cfg.KubeContext,
		})

	clientConfig, err := config.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("can't build config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("can't build client: %w", err)
	}

	if cfg.Namespace == "" {
		cfg.Namespace, _, err = config.Namespace()
		if err != nil {
			return nil, fmt.Errorf("can't get current namespace: %w", err)

		}
	}

	return clientset, nil
}

func getDynamicClient(cfg *config.Config) (*dynamic.DynamicClient, error) {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{Precedence: strings.Split(cfg.KubeConfig, ":")},
		&clientcmd.ConfigOverrides{
			CurrentContext: cfg.KubeContext,
		})

	clientConfig, err := config.ClientConfig()

	dynamicClient, err := dynamic.NewForConfig(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("can't build dynamic client: %w", err)
	}

	return dynamicClient, nil
}
