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
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kmodules.xyz/resource-metadata/apis/ui/v1alpha1"
)

// FakeResourceDashboards implements ResourceDashboardInterface
type FakeResourceDashboards struct {
	Fake *FakeUiV1alpha1
}

var resourcedashboardsResource = schema.GroupVersionResource{Group: "ui.k8s.appscode.com", Version: "v1alpha1", Resource: "resourcedashboards"}

var resourcedashboardsKind = schema.GroupVersionKind{Group: "ui.k8s.appscode.com", Version: "v1alpha1", Kind: "ResourceDashboard"}

// Get takes name of the resourceDashboard, and returns the corresponding resourceDashboard object, and an error if there is any.
func (c *FakeResourceDashboards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(resourcedashboardsResource, name), &v1alpha1.ResourceDashboard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDashboard), err
}

// List takes label and field selectors, and returns the list of ResourceDashboards that match those selectors.
func (c *FakeResourceDashboards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceDashboardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(resourcedashboardsResource, resourcedashboardsKind, opts), &v1alpha1.ResourceDashboardList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceDashboardList{ListMeta: obj.(*v1alpha1.ResourceDashboardList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceDashboardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceDashboards.
func (c *FakeResourceDashboards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(resourcedashboardsResource, opts))
}

// Create takes the representation of a resourceDashboard and creates it.  Returns the server's representation of the resourceDashboard, and an error, if there is any.
func (c *FakeResourceDashboards) Create(ctx context.Context, resourceDashboard *v1alpha1.ResourceDashboard, opts v1.CreateOptions) (result *v1alpha1.ResourceDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(resourcedashboardsResource, resourceDashboard), &v1alpha1.ResourceDashboard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDashboard), err
}

// Update takes the representation of a resourceDashboard and updates it. Returns the server's representation of the resourceDashboard, and an error, if there is any.
func (c *FakeResourceDashboards) Update(ctx context.Context, resourceDashboard *v1alpha1.ResourceDashboard, opts v1.UpdateOptions) (result *v1alpha1.ResourceDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(resourcedashboardsResource, resourceDashboard), &v1alpha1.ResourceDashboard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDashboard), err
}

// Delete takes name of the resourceDashboard and deletes it. Returns an error if one occurs.
func (c *FakeResourceDashboards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(resourcedashboardsResource, name, opts), &v1alpha1.ResourceDashboard{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceDashboards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(resourcedashboardsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceDashboardList{})
	return err
}

// Patch applies the patch and returns the patched resourceDashboard.
func (c *FakeResourceDashboards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceDashboard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(resourcedashboardsResource, name, pt, data, subresources...), &v1alpha1.ResourceDashboard{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDashboard), err
}
