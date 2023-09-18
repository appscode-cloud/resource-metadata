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

package v1alpha1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	scheme "kmodules.xyz/resource-metadata/client/clientset/versioned/scheme"
)

// RenderDashboardsGetter has a method to return a RenderDashboardInterface.
// A group's client should implement this interface.
type RenderDashboardsGetter interface {
	RenderDashboards() RenderDashboardInterface
}

// RenderDashboardInterface has methods to work with RenderDashboard resources.
type RenderDashboardInterface interface {
	Create(ctx context.Context, renderDashboard *v1alpha1.RenderDashboard, opts v1.CreateOptions) (*v1alpha1.RenderDashboard, error)
	RenderDashboardExpansion
}

// renderDashboards implements RenderDashboardInterface
type renderDashboards struct {
	client rest.Interface
}

// newRenderDashboards returns a RenderDashboards
func newRenderDashboards(c *MetaV1alpha1Client) *renderDashboards {
	return &renderDashboards{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a renderDashboard and creates it.  Returns the server's representation of the renderDashboard, and an error, if there is any.
func (c *renderDashboards) Create(ctx context.Context, renderDashboard *v1alpha1.RenderDashboard, opts v1.CreateOptions) (result *v1alpha1.RenderDashboard, err error) {
	result = &v1alpha1.RenderDashboard{}
	err = c.client.Post().
		Resource("renderdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(renderDashboard).
		Do(ctx).
		Into(result)
	return
}
