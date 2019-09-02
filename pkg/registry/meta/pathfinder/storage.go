package pathfinder

import (
	"context"

	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	"kmodules.xyz/resource-metadata/pkg/graph"
)

type Storage struct {
}

var _ rest.GroupVersionKindProvider = &Storage{}
var _ rest.Scoper = &Storage{}
var _ rest.Creater = &Storage{}

func NewStorage() *Storage {
	return &Storage{}
}

func (r *Storage) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
	return v1alpha1.SchemeGroupVersion.WithKind(v1alpha1.ResourceKindPathFinder)
}

func (r *Storage) NamespaceScoped() bool {
	return false
}

// Getter
func (r *Storage) New() runtime.Object {
	return &v1alpha1.PathFinder{}
}

func (r *Storage) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	pf := obj.(*v1alpha1.PathFinder)

	srcGVR := schema.GroupVersionResource{
		Group:    pf.Request.Source.Group,
		Version:  pf.Request.Source.Version,
		Resource: pf.Request.Source.Resource,
	}

	var targetGVR *schema.GroupVersionResource
	if pf.Request.Target != nil {
		targetGVR = &schema.GroupVersionResource{
			Group:    pf.Request.Target.Group,
			Version:  pf.Request.Target.Version,
			Resource: pf.Request.Target.Resource,
		}
	}

	g, err := graph.LoadGraph()
	if err != nil {
		return nil, kerr.NewInternalError(err)
	}

	dist, prev := graph.Dijkstra(g, srcGVR)
	paths := graph.GeneratePaths(srcGVR, dist, prev)

	out := make([]v1alpha1.Path, 0, len(paths))

	if targetGVR != nil {
		path, ok := paths[*targetGVR]
		if ok {
			out = append(out, convertPath(*path))
		}
	} else {
		for i := range paths {
			out = append(out, convertPath(*paths[i]))
		}
	}

	pf.Response = &v1alpha1.PathResponse{Paths: out}
	return pf, nil
}

func convertPath(in graph.Path) v1alpha1.Path {
	out := v1alpha1.Path{
		Source: v1alpha1.GroupVersionResource{
			Group:    in.Source.Group,
			Version:  in.Source.Version,
			Resource: in.Source.Resource,
		},
		Target: v1alpha1.GroupVersionResource{
			Group:    in.Target.Group,
			Version:  in.Target.Version,
			Resource: in.Target.Resource,
		},
		Distance: in.Distance,
		Edges:    make([]v1alpha1.Edge, len(in.Edges)),
	}

	for i := range in.Edges {
		out.Edges[i] = convertEdge(*in.Edges[i])
	}

	return out
}

func convertEdge(in graph.Edge) v1alpha1.Edge {
	return v1alpha1.Edge{
		Src: v1alpha1.GroupVersionResource{
			Group:    in.Src.Group,
			Version:  in.Src.Version,
			Resource: in.Src.Resource,
		},
		Dst: v1alpha1.GroupVersionResource{
			Group:    in.Dst.Group,
			Version:  in.Dst.Version,
			Resource: in.Dst.Resource,
		},
		W:          in.W,
		Connection: in.Connection,
		Forward:    in.Forward,
	}
}
