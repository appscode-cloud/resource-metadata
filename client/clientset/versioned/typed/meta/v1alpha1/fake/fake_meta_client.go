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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kmodules.xyz/resource-metadata/client/clientset/versioned/typed/meta/v1alpha1"
)

type FakeMetaV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMetaV1alpha1) ChartPresetQueries() v1alpha1.ChartPresetQueryInterface {
	return &FakeChartPresetQueries{c}
}

func (c *FakeMetaV1alpha1) ClusterStatuses() v1alpha1.ClusterStatusInterface {
	return &FakeClusterStatuses{c}
}

func (c *FakeMetaV1alpha1) MenuOutlines(namespace string) v1alpha1.MenuOutlineInterface {
	return &FakeMenuOutlines{c, namespace}
}

func (c *FakeMetaV1alpha1) Renders() v1alpha1.RenderInterface {
	return &FakeRenders{c}
}

func (c *FakeMetaV1alpha1) RenderDashboards() v1alpha1.RenderDashboardInterface {
	return &FakeRenderDashboards{c}
}

func (c *FakeMetaV1alpha1) RenderMenus() v1alpha1.RenderMenuInterface {
	return &FakeRenderMenus{c}
}

func (c *FakeMetaV1alpha1) RenderRawGraphs() v1alpha1.RenderRawGraphInterface {
	return &FakeRenderRawGraphs{c}
}

func (c *FakeMetaV1alpha1) ResourceBlockDefinitions() v1alpha1.ResourceBlockDefinitionInterface {
	return &FakeResourceBlockDefinitions{c}
}

func (c *FakeMetaV1alpha1) ResourceCalculators() v1alpha1.ResourceCalculatorInterface {
	return &FakeResourceCalculators{c}
}

func (c *FakeMetaV1alpha1) ResourceDescriptors() v1alpha1.ResourceDescriptorInterface {
	return &FakeResourceDescriptors{c}
}

func (c *FakeMetaV1alpha1) ResourceEditors() v1alpha1.ResourceEditorInterface {
	return &FakeResourceEditors{c}
}

func (c *FakeMetaV1alpha1) ResourceGraphs() v1alpha1.ResourceGraphInterface {
	return &FakeResourceGraphs{c}
}

func (c *FakeMetaV1alpha1) ResourceLayouts() v1alpha1.ResourceLayoutInterface {
	return &FakeResourceLayouts{c}
}

func (c *FakeMetaV1alpha1) ResourceOutlines() v1alpha1.ResourceOutlineInterface {
	return &FakeResourceOutlines{c}
}

func (c *FakeMetaV1alpha1) ResourceQueries() v1alpha1.ResourceQueryInterface {
	return &FakeResourceQueries{c}
}

func (c *FakeMetaV1alpha1) ResourceTableDefinitions() v1alpha1.ResourceTableDefinitionInterface {
	return &FakeResourceTableDefinitions{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMetaV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
