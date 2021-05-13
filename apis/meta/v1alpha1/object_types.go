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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ObjectLocator struct {
	Start *ObjectRef `json:"start"`
	Path  []string   `json:"path"` // sequence of DirectedEdge names
}

type NamedEdge struct {
	Name       string                 `json:"name"`
	Src        metav1.TypeMeta        `json:"src"`
	Dst        metav1.TypeMeta        `json:"dst"`
	Connection ResourceConnectionSpec `json:"connection"`
}

// Namespace always same as Workflow
type ObjectRef struct {
	Target       metav1.TypeMeta       `json:"target"`
	Selector     *metav1.LabelSelector `json:"selector,omitempty"`
	NameTemplate string                `json:"nameTemplate,omitempty"`
}
