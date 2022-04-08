package pdbs

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// The k8s.io/client-go/plugin/pkg/client/auth import ensures that client-go can authenticate with Kubernetes clusters on cloud providers.
	// See https://krew.sigs.k8s.io/docs/developer-guide/develop/best-practices/
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

// GetBorkedPdbs will return any pod disruption budgets that currently have zero disruptions allowed
// namespace: The namespace you wish to filter on. This can be left empty, in which case the search will span the entire cluster.
func GetBorkedPdbs(namespace string) error {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{})

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	results, err := clientset.PolicyV1().PodDisruptionBudgets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pdb := range results.Items {
		if pdb.Status.DisruptionsAllowed == 0 {
			fmt.Println(pdb.Name)
		}
	}

	return nil
}
