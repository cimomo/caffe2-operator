/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	caffe2_v1alpha1 "github.com/kubeflow/caffe2-operator/pkg/apis/caffe2/v1alpha1"
	versioned "github.com/kubeflow/caffe2-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/caffe2-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeflow/caffe2-operator/pkg/client/listers/kubeflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// Caffe2JobInformer provides access to a shared informer and lister for
// Caffe2Jobs.
type Caffe2JobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.Caffe2JobLister
}

type caffe2JobInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewCaffe2JobInformer constructs a new informer for Caffe2Job type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCaffe2JobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.KubeflowV1alpha1().Caffe2Jobs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.KubeflowV1alpha1().Caffe2Jobs(namespace).Watch(options)
			},
		},
		&caffe2_v1alpha1.Caffe2Job{},
		resyncPeriod,
		indexers,
	)
}

func defaultCaffe2JobInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewCaffe2JobInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *caffe2JobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&caffe2_v1alpha1.Caffe2Job{}, defaultCaffe2JobInformer)
}

func (f *caffe2JobInformer) Lister() v1alpha1.Caffe2JobLister {
	return v1alpha1.NewCaffe2JobLister(f.Informer().GetIndexer())
}