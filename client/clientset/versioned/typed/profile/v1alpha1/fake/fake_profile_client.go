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
	v1alpha1 "kmodules.xyz/resource-metadata/client/clientset/versioned/typed/profile/v1alpha1"
)

type FakeProfileV1alpha1 struct {
	*testing.Fake
}

func (c *FakeProfileV1alpha1) ClusterProfiles() v1alpha1.ClusterProfileInterface {
	return &FakeClusterProfiles{c}
}

func (c *FakeProfileV1alpha1) ManagedClusterProfileBindings(namespace string) v1alpha1.ManagedClusterProfileBindingInterface {
	return &FakeManagedClusterProfileBindings{c, namespace}
}

func (c *FakeProfileV1alpha1) ManagedClusterSetProfiles() v1alpha1.ManagedClusterSetProfileInterface {
	return &FakeManagedClusterSetProfiles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProfileV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
