/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
)

// FakeResourceLayouts implements ResourceLayoutInterface
type FakeResourceLayouts struct {
	Fake *FakeMetaV1alpha1
}

var resourcelayoutsResource = schema.GroupVersionResource{Group: "meta.k8s.appscode.com", Version: "v1alpha1", Resource: "resourcelayouts"}

var resourcelayoutsKind = schema.GroupVersionKind{Group: "meta.k8s.appscode.com", Version: "v1alpha1", Kind: "ResourceLayout"}

// Get takes name of the resourceLayout, and returns the corresponding resourceLayout object, and an error if there is any.
func (c *FakeResourceLayouts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceLayout, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(resourcelayoutsResource, name), &v1alpha1.ResourceLayout{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceLayout), err
}
