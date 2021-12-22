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

import (
	"fmt"
	"strings"

	apiv1 "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/apiextensions"
	"kmodules.xyz/resource-metadata/crds"
)

func (v ResourceDescriptor) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceResourceDescriptors))
}

func (v ResourceDescriptor) IsValid() error {
	return nil
}

func IsOfficialType(group string) bool {
	switch {
	case group == "":
		return true
	case !strings.ContainsRune(group, '.'):
		return true
	case group == "k8s.io" || strings.HasSuffix(group, ".k8s.io"):
		return true
	case group == "kubernetes.io" || strings.HasSuffix(group, ".kubernetes.io"):
		return true
	case group == "x-k8s.io" || strings.HasSuffix(group, ".x-k8s.io"):
		return true
	default:
		return false
	}
}

func (r ResourceLocator) GraphQuery(oid apiv1.OID) (string, map[string]interface{}) {
	vars := map[string]interface{}{
		"src":         string(oid),
		"targetGroup": r.Ref.Group,
		"targetKind":  r.Ref.Kind,
	}

	if r.Query.Raw != "" {
		return r.Query.Raw, vars
	}
	return fmt.Sprintf(`query Find($src: String!, $targetGroup: String!, $targetKind: String!) {
  find(oid: $src) {
    refs: %s(group: $targetGroup, kind: $targetKind) {
      namespace
      name
    }
  }
}`, r.Query.ByLabel), vars
}
