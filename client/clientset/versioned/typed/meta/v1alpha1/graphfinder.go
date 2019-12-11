/*
Copyright The Kmodules Authors.

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
	rest "k8s.io/client-go/rest"
	v1alpha1 "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
)

// GraphFindersGetter has a method to return a GraphFinderInterface.
// A group's client should implement this interface.
type GraphFindersGetter interface {
	GraphFinders() GraphFinderInterface
}

// GraphFinderInterface has methods to work with GraphFinder resources.
type GraphFinderInterface interface {
	Create(*v1alpha1.GraphFinder) (*v1alpha1.GraphFinder, error)
	GraphFinderExpansion
}

// graphFinders implements GraphFinderInterface
type graphFinders struct {
	client rest.Interface
}

// newGraphFinders returns a GraphFinders
func newGraphFinders(c *MetaV1alpha1Client) *graphFinders {
	return &graphFinders{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a graphFinder and creates it.  Returns the server's representation of the graphFinder, and an error, if there is any.
func (c *graphFinders) Create(graphFinder *v1alpha1.GraphFinder) (result *v1alpha1.GraphFinder, err error) {
	result = &v1alpha1.GraphFinder{}
	err = c.client.Post().
		Resource("graphfinders").
		Body(graphFinder).
		Do().
		Into(result)
	return
}
