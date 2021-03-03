package clusterapi

import "k8s.io/client-go/kubernetes"

type Cluster struct {
	clientset *kubernetes.Clientset
	Pods      []*Pods
	Svcs      []*Services
	Pv        int
}
type Pods struct {
	Name       string
	Namespace  string
	CPURequest int
	MemRequest int
	CPULimit   int
	MemLimit   int
}
type Services struct {
	Name       string
	Namespace  string
	CPURequest int
	MemRequest int
	CPULimit   int
	MemLimit   int
}
