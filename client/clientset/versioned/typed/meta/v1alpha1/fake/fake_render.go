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

// FakeRenders implements RenderInterface
type FakeRenders struct {
	Fake *FakeMetaV1alpha1
}

var rendersResource = schema.GroupVersionResource{Group: "meta.k8s.appscode.com", Version: "v1alpha1", Resource: "renders"}

var rendersKind = schema.GroupVersionKind{Group: "meta.k8s.appscode.com", Version: "v1alpha1", Kind: "Render"}

// Create takes the representation of a render and creates it.  Returns the server's representation of the render, and an error, if there is any.
func (c *FakeRenders) Create(ctx context.Context, render *v1alpha1.Render, opts v1.CreateOptions) (result *v1alpha1.Render, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rendersResource, render), &v1alpha1.Render{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Render), err
}
