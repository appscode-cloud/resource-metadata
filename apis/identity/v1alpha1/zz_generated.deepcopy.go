//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/authorization/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIdentity) DeepCopyInto(out *ClusterIdentity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIdentity.
func (in *ClusterIdentity) DeepCopy() *ClusterIdentity {
	if in == nil {
		return nil
	}
	out := new(ClusterIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIdentity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIdentityList) DeepCopyInto(out *ClusterIdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIdentityList.
func (in *ClusterIdentityList) DeepCopy() *ClusterIdentityList {
	if in == nil {
		return nil
	}
	out := new(ClusterIdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIdentityStatus) DeepCopyInto(out *ClusterIdentityStatus) {
	*out = *in
	out.ClusterMetadata = in.ClusterMetadata
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIdentityStatus.
func (in *ClusterIdentityStatus) DeepCopy() *ClusterIdentityStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterIdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboxTokenRequest) DeepCopyInto(out *InboxTokenRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(InboxTokenRequestRequest)
		**out = **in
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(InboxTokenRequestResponse)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboxTokenRequest.
func (in *InboxTokenRequest) DeepCopy() *InboxTokenRequest {
	if in == nil {
		return nil
	}
	out := new(InboxTokenRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InboxTokenRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboxTokenRequestRequest) DeepCopyInto(out *InboxTokenRequestRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboxTokenRequestRequest.
func (in *InboxTokenRequestRequest) DeepCopy() *InboxTokenRequestRequest {
	if in == nil {
		return nil
	}
	out := new(InboxTokenRequestRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboxTokenRequestResponse) DeepCopyInto(out *InboxTokenRequestResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboxTokenRequestResponse.
func (in *InboxTokenRequestResponse) DeepCopy() *InboxTokenRequestResponse {
	if in == nil {
		return nil
	}
	out := new(InboxTokenRequestResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectNamespaceAccessReview) DeepCopyInto(out *SelfSubjectNamespaceAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectNamespaceAccessReview.
func (in *SelfSubjectNamespaceAccessReview) DeepCopy() *SelfSubjectNamespaceAccessReview {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectNamespaceAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectNamespaceAccessReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectNamespaceAccessReviewList) DeepCopyInto(out *SelfSubjectNamespaceAccessReviewList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SelfSubjectNamespaceAccessReview, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectNamespaceAccessReviewList.
func (in *SelfSubjectNamespaceAccessReviewList) DeepCopy() *SelfSubjectNamespaceAccessReviewList {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectNamespaceAccessReviewList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectNamespaceAccessReviewList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectNamespaceAccessReviewSpec) DeepCopyInto(out *SelfSubjectNamespaceAccessReviewSpec) {
	*out = *in
	if in.ResourceAttributes != nil {
		in, out := &in.ResourceAttributes, &out.ResourceAttributes
		*out = make([]v1.ResourceAttributes, len(*in))
		copy(*out, *in)
	}
	if in.NonResourceAttributes != nil {
		in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
		*out = make([]v1.NonResourceAttributes, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectNamespaceAccessReviewSpec.
func (in *SelfSubjectNamespaceAccessReviewSpec) DeepCopy() *SelfSubjectNamespaceAccessReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectNamespaceAccessReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessNamespaceReviewStatus) DeepCopyInto(out *SubjectAccessNamespaceReviewStatus) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessNamespaceReviewStatus.
func (in *SubjectAccessNamespaceReviewStatus) DeepCopy() *SubjectAccessNamespaceReviewStatus {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessNamespaceReviewStatus)
	in.DeepCopyInto(out)
	return out
}
