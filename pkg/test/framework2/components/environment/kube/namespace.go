//  Copyright 2019 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package kube

import (
	"io"

	"istio.io/istio/pkg/test/framework2/core"

	"istio.io/istio/pkg/test/kube"
)

// kubeNamespace represents a Kubernetes namespace. It is tracked as a resource.
type kubeNamespace struct {
	id   core.ResourceID
	name string
	a    *kube.Accessor
}

var _ core.Namespace = &kubeNamespace{}
var _ io.Closer = &kubeNamespace{}
var _ core.Resource = &kubeNamespace{}
var _ core.Dumper = &kubeNamespace{}

func (n *kubeNamespace) Name() string {
	return n.name
}

func (n *kubeNamespace) ID() core.ResourceID {
	return n.id
}

// Close implements io.Closer
func (n *kubeNamespace) Close() error {
	if n.name != "" {
		ns := n.name
		n.name = ""
		return n.a.DeleteNamespace(ns)
	}

	return nil
}

// Dump implements resource.Dumper
func (n *kubeNamespace) Dump() {
	// TODO: Make this dumpable.
}
