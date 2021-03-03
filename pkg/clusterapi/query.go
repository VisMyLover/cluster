package clusterapi

import (
	"context"
	"errors"

	"github.com/prometheus/common/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type pvInfo struct {
	PvName    string
	NameSpace string
	PvcName   string
}

var pvs []*pvInfo

func (c *Cluster) queryClusterPods() {
	ctx, _ := context.WithCancel(context.Background())
	pods, err := c.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error("Fail to query pods, err: ", err)
	}

	for _, pod := range pods.Items {
		p := new(Pods)
		p.Name = pod.Name
		p.Namespace = pod.Namespace
		c.Pods = append(c.Pods, p)
	}
}

func (c *Cluster) queryClusterServices() {
	ctx, _ := context.WithCancel(context.Background())
	services, err := c.clientset.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error("Fail to query pods, err: ", err)
	}
	for _, svc := range services.Items {
		s := new(Services)
		s.Name = svc.Name
		s.Namespace = svc.Namespace
		c.Svcs = append(c.Svcs, s)
	}
}

func (c *Cluster) queryClusterPvs() {
	ctx, _ := context.WithCancel(context.Background())
	pv, err := c.clientset.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	checkError(err)
	c.Pv = len(pv.Items)
	pvs, err = c.fifterPvInfo(pv)
}
func (c *Cluster) fifterPvInfo(p *v1.PersistentVolumeList) ([]*pvInfo, error) {
	if len(p.Items) == 0 {
		return nil, errors.New("pv items is nil")
	}

	pvs2 := make([]*pvInfo, 0)

	log.Info("PV num: ", len(p.Items))

	for _, item := range p.Items {
		if item.Spec.ClaimRef == nil {
			continue
		}

		pv := new(pvInfo)
		pv.PvName = item.Name
		pv.PvcName = item.Spec.ClaimRef.Name
		pv.NameSpace = item.Spec.ClaimRef.Namespace
		pvs2 = append(pvs2, pv)
	}
	return pvs2, nil
}

func checkError(err error) {
	if err != nil {
		log.Error(err)
	}
}
