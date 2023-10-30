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

package shared

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "x-helm.dev/apimachinery/apis/releases/v1alpha1"
	apisshared "x-helm.dev/apimachinery/apis/shared"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]apisshared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionGroup) DeepCopyInto(out *ActionGroup) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Action, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionGroup.
func (in *ActionGroup) DeepCopy() *ActionGroup {
	if in == nil {
		return nil
	}
	out := new(ActionGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionInfo) DeepCopyInto(out *ActionInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionInfo.
func (in *ActionInfo) DeepCopy() *ActionInfo {
	if in == nil {
		return nil
	}
	out := new(ActionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionTemplate) DeepCopyInto(out *ActionTemplate) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]apisshared.ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionTemplate.
func (in *ActionTemplate) DeepCopy() *ActionTemplate {
	if in == nil {
		return nil
	}
	out := new(ActionTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionTemplateGroup) DeepCopyInto(out *ActionTemplateGroup) {
	*out = *in
	out.ActionInfo = in.ActionInfo
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActionTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionTemplateGroup.
func (in *ActionTemplateGroup) DeepCopy() *ActionTemplateGroup {
	if in == nil {
		return nil
	}
	out := new(ActionTemplateGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapPresets) DeepCopyInto(out *BootstrapPresets) {
	*out = *in
	out.Image = in.Image
	in.Registry.DeepCopyInto(&out.Registry)
	in.Helm.DeepCopyInto(&out.Helm)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapPresets.
func (in *BootstrapPresets) DeepCopy() *BootstrapPresets {
	if in == nil {
		return nil
	}
	out := new(BootstrapPresets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]DashboardVar, len(*in))
		copy(*out, *in)
	}
	if in.Panels != nil {
		in, out := &in.Panels, &out.Panels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.If != nil {
		in, out := &in.If, &out.If
		*out = new(If)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardVar) DeepCopyInto(out *DashboardVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardVar.
func (in *DashboardVar) DeepCopy() *DashboardVar {
	if in == nil {
		return nil
	}
	out := new(DashboardVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentParameters) DeepCopyInto(out *DeploymentParameters) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentParameters.
func (in *DeploymentParameters) DeepCopy() *DeploymentParameters {
	if in == nil {
		return nil
	}
	out := new(DeploymentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInfo) DeepCopyInto(out *HelmInfo) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make(map[string]*HelmRepository, len(*in))
		for key, val := range *in {
			var outVal *HelmRepository
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(HelmRepository)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Releases != nil {
		in, out := &in.Releases, &out.Releases
		*out = make(map[string]*HelmRelease, len(*in))
		for key, val := range *in {
			var outVal *HelmRelease
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(HelmRelease)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInfo.
func (in *HelmInfo) DeepCopy() *HelmInfo {
	if in == nil {
		return nil
	}
	out := new(HelmInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRelease) DeepCopyInto(out *HelmRelease) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRelease.
func (in *HelmRelease) DeepCopy() *HelmRelease {
	if in == nil {
		return nil
	}
	out := new(HelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRepository) DeepCopyInto(out *HelmRepository) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRepository.
func (in *HelmRepository) DeepCopy() *HelmRepository {
	if in == nil {
		return nil
	}
	out := new(HelmRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *If) DeepCopyInto(out *If) {
	*out = *in
	if in.Connected != nil {
		in, out := &in.Connected, &out.Connected
		*out = new(ResourceLocator)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new If.
func (in *If) DeepCopy() *If {
	if in == nil {
		return nil
	}
	out := new(If)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRegistrySpec) DeepCopyInto(out *ImageRegistrySpec) {
	*out = *in
	out.Proxies = in.Proxies
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRegistrySpec.
func (in *ImageRegistrySpec) DeepCopy() *ImageRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ImageRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryInfo) DeepCopyInto(out *RegistryInfo) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make(map[string]Credentials, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryInfo.
func (in *RegistryInfo) DeepCopy() *RegistryInfo {
	if in == nil {
		return nil
	}
	out := new(RegistryInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryProxies) DeepCopyInto(out *RegistryProxies) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryProxies.
func (in *RegistryProxies) DeepCopy() *RegistryProxies {
	if in == nil {
		return nil
	}
	out := new(RegistryProxies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLocator) DeepCopyInto(out *ResourceLocator) {
	*out = *in
	out.Ref = in.Ref
	out.Query = in.Query
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLocator.
func (in *ResourceLocator) DeepCopy() *ResourceLocator {
	if in == nil {
		return nil
	}
	out := new(ResourceLocator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuery) DeepCopyInto(out *ResourceQuery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuery.
func (in *ResourceQuery) DeepCopy() *ResourceQuery {
	if in == nil {
		return nil
	}
	out := new(ResourceQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceLocator) DeepCopyInto(out *SourceLocator) {
	*out = *in
	out.Resource = in.Resource
	out.Ref = in.Ref
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceLocator.
func (in *SourceLocator) DeepCopy() *SourceLocator {
	if in == nil {
		return nil
	}
	out := new(SourceLocator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UIParameterTemplate) DeepCopyInto(out *UIParameterTemplate) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*ActionTemplateGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActionTemplateGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.InstanceLabelPaths != nil {
		in, out := &in.InstanceLabelPaths, &out.InstanceLabelPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UIParameterTemplate.
func (in *UIParameterTemplate) DeepCopy() *UIParameterTemplate {
	if in == nil {
		return nil
	}
	out := new(UIParameterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UIParameters) DeepCopyInto(out *UIParameters) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.Editor != nil {
		in, out := &in.Editor, &out.Editor
		*out = new(v1alpha1.ChartSourceRef)
		**out = **in
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*ActionGroup, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActionGroup)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.InstanceLabelPaths != nil {
		in, out := &in.InstanceLabelPaths, &out.InstanceLabelPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UIParameters.
func (in *UIParameters) DeepCopy() *UIParameters {
	if in == nil {
		return nil
	}
	out := new(UIParameters)
	in.DeepCopyInto(out)
	return out
}
