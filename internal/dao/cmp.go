// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package dao

import (
	"errors"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	_ Accessor = (*Cluster)(nil)
	_ Nuker    = (*Cluster)(nil)
)

// Cluster represents a K8s sts.
type Cluster struct {
	Resource
}

// GetInstance returns a statefulset instance.
func (*Cluster) GetInstance(f Factory, fqn string) (*appsv1.Cluster, error) {
	o, err := f.Get("apps.kubeblocks.io/v1/clusters", fqn, true, labels.Everything())
	if err != nil {
		return nil, err
	}

	var cluster appsv1.Cluster
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &cluster)
	if err != nil {
		return nil, errors.New("expecting Statefulset resource")
	}

	return &cluster, nil
}
