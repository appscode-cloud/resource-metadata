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
	v1alpha1 "kmodules.xyz/resource-metadata/apis/identity/v1alpha1"
	scheme "kmodules.xyz/resource-metadata/client/clientset/versioned/scheme"
)

// InboxTokenRequestsGetter has a method to return a InboxTokenRequestInterface.
// A group's client should implement this interface.
type InboxTokenRequestsGetter interface {
	InboxTokenRequests() InboxTokenRequestInterface
}

// InboxTokenRequestInterface has methods to work with InboxTokenRequest resources.
type InboxTokenRequestInterface interface {
	Create(ctx context.Context, inboxTokenRequest *v1alpha1.InboxTokenRequest, opts v1.CreateOptions) (*v1alpha1.InboxTokenRequest, error)
	InboxTokenRequestExpansion
}

// inboxTokenRequests implements InboxTokenRequestInterface
type inboxTokenRequests struct {
	client rest.Interface
}

// newInboxTokenRequests returns a InboxTokenRequests
func newInboxTokenRequests(c *IdentityV1alpha1Client) *inboxTokenRequests {
	return &inboxTokenRequests{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a inboxTokenRequest and creates it.  Returns the server's representation of the inboxTokenRequest, and an error, if there is any.
func (c *inboxTokenRequests) Create(ctx context.Context, inboxTokenRequest *v1alpha1.InboxTokenRequest, opts v1.CreateOptions) (result *v1alpha1.InboxTokenRequest, err error) {
	result = &v1alpha1.InboxTokenRequest{}
	err = c.client.Post().
		Resource("inboxtokenrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inboxTokenRequest).
		Do(ctx).
		Into(result)
	return
}
