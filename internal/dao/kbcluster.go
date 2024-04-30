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
	_ Accessor = (*Component)(nil)
	_ Nuker    = (*Component)(nil)
)

// Component represents a K8s sts.
type Component struct {
	Resource
}

// GetInstance returns a statefulset instance.
func (*Component) GetInstance(f Factory, fqn string) (*appsv1.Component, error) {
	o, err := f.Get("apps.kubeblocks.io/v1/components", fqn, true, labels.Everything())
	if err != nil {
		return nil, err
	}

	var cmp appsv1.Component
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &cmp)
	if err != nil {
		return nil, errors.New("expecting Statefulset resource")
	}

	return &cmp, nil
}
