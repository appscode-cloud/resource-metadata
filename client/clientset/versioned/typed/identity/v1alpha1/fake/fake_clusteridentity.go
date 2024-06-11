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
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kmodules.xyz/resource-metadata/apis/identity/v1alpha1"
)

// FakeClusterIdentities implements ClusterIdentityInterface
type FakeClusterIdentities struct {
	Fake *FakeIdentityV1alpha1
	ns   string
}

var clusteridentitiesResource = v1alpha1.SchemeGroupVersion.WithResource("clusteridentities")

var clusteridentitiesKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterIdentity")

// Get takes name of the clusterIdentity, and returns the corresponding clusterIdentity object, and an error if there is any.
func (c *FakeClusterIdentities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteridentitiesResource, c.ns, name), &v1alpha1.ClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIdentity), err
}

// List takes label and field selectors, and returns the list of ClusterIdentities that match those selectors.
func (c *FakeClusterIdentities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterIdentityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteridentitiesResource, clusteridentitiesKind, c.ns, opts), &v1alpha1.ClusterIdentityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterIdentityList{ListMeta: obj.(*v1alpha1.ClusterIdentityList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterIdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterIdentities.
func (c *FakeClusterIdentities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteridentitiesResource, c.ns, opts))

}

// Create takes the representation of a clusterIdentity and creates it.  Returns the server's representation of the clusterIdentity, and an error, if there is any.
func (c *FakeClusterIdentities) Create(ctx context.Context, clusterIdentity *v1alpha1.ClusterIdentity, opts v1.CreateOptions) (result *v1alpha1.ClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteridentitiesResource, c.ns, clusterIdentity), &v1alpha1.ClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIdentity), err
}

// Update takes the representation of a clusterIdentity and updates it. Returns the server's representation of the clusterIdentity, and an error, if there is any.
func (c *FakeClusterIdentities) Update(ctx context.Context, clusterIdentity *v1alpha1.ClusterIdentity, opts v1.UpdateOptions) (result *v1alpha1.ClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteridentitiesResource, c.ns, clusterIdentity), &v1alpha1.ClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIdentity), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterIdentities) UpdateStatus(ctx context.Context, clusterIdentity *v1alpha1.ClusterIdentity, opts v1.UpdateOptions) (*v1alpha1.ClusterIdentity, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusteridentitiesResource, "status", c.ns, clusterIdentity), &v1alpha1.ClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIdentity), err
}

// Delete takes name of the clusterIdentity and deletes it. Returns an error if one occurs.
func (c *FakeClusterIdentities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusteridentitiesResource, c.ns, name, opts), &v1alpha1.ClusterIdentity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterIdentities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteridentitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterIdentityList{})
	return err
}

// Patch applies the patch and returns the patched clusterIdentity.
func (c *FakeClusterIdentities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteridentitiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIdentity), err
}
