package clusterapi

func New() *Cluster {
	return new(Cluster)
}

func Count() {
	c := New()
	c.clientset = c.informer()
	c.queryClusterPods()
	c.queryClusterServices()
}
