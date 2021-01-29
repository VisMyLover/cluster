package clusterapi

import (
	"flag"

	"github.com/prometheus/common/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func (c *Cluster) informer() *kubernetes.Clientset {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Error(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err)
	}

	return clientset
}
