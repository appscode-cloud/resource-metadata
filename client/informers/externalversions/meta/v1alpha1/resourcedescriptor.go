/*
Copyright 2019 The ResourceMetadata Project Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	metav1alpha1 "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	versioned "kmodules.xyz/resource-metadata/client/clientset/versioned"
	internalinterfaces "kmodules.xyz/resource-metadata/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kmodules.xyz/resource-metadata/client/listers/meta/v1alpha1"
)

// ResourceDescriptorInformer provides access to a shared informer and lister for
// ResourceDescriptors.
type ResourceDescriptorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceDescriptorLister
}

type resourceDescriptorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceDescriptorInformer constructs a new informer for ResourceDescriptor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceDescriptorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceDescriptorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceDescriptorInformer constructs a new informer for ResourceDescriptor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceDescriptorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().ResourceDescriptors().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MetaV1alpha1().ResourceDescriptors().Watch(options)
			},
		},
		&metav1alpha1.ResourceDescriptor{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceDescriptorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceDescriptorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceDescriptorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&metav1alpha1.ResourceDescriptor{}, f.defaultInformer)
}

func (f *resourceDescriptorInformer) Lister() v1alpha1.ResourceDescriptorLister {
	return v1alpha1.NewResourceDescriptorLister(f.Informer().GetIndexer())
}
