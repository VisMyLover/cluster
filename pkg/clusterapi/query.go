package clusterapi

import (
	"context"

	"github.com/prometheus/common/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Cluster) queryClusterPods() {
	ctx, _ := context.WithCancel(context.Background())
	pods, err := c.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error("Fail to query pods, err: ", err)
	}
	log.Info("pods len is ", len(pods.Items))

	for _, pod := range pods.Items {
		p := new(Pods)
		p.Name = pod.Name
		p.Namespace = pod.Namespace
		c.Pods = append(c.Pods, p)
		//for _, container := range pod.Spec.Containers {
		//	p.CPURequest += container.Resources.Requests["cpu"].Quantity
		//}
		log.Info(pod.Name, " ", pod.Namespace, " ", pod.ClusterName)
	}
}

func (c *Cluster) queryClusterServices() {
	ctx, _ := context.WithCancel(context.Background())
	services, err := c.clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error("Fail to query pods, err: ", err)
	}
	log.Info("pods len is ", len(services.Items))
	for _, svc := range services.Items {
		s := new(Services)
		s.Name = svc.Name
		s.Namespace = svc.Namespace
		c.Svcs = append(c.Svcs, s)
		log.Info(svc.Name, " ", svc.Namespace, " ", svc.ClusterName)
	}
}
