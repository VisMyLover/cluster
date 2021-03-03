package clusterapi

import "github.com/prometheus/common/log"

func New() *Cluster {
	return new(Cluster)
}

func Count() {
	c := New()
	c.clientset = c.informer()
	c.queryClusterPods()
	c.queryClusterServices()
	c.queryClusterPvs()

	c.inPutClusterInto()

}
func (c *Cluster) inPutClusterInto() {
	log.Info("pods: ", len(c.Pods))
	log.Info("services: ", len(c.Svcs))
	log.Info("pvs: ", c.Pv)
}
