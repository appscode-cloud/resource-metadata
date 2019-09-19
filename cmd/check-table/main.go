package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/resource-metadata/pkg/tableconvertor"
)

func main() {
	masterURL := ""
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}

	dc, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	gvr := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	{
		list, err := dc.Resource(gvr).Namespace("default").List(metav1.ListOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		t, err := tableconvertor.TableForList(gvr, list.Items)
		if err != nil {
			log.Fatalln(err)
		}

		data, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(data))
	}

	{
		dep, err := dc.Resource(gvr).Namespace("default").Get("busy-dep", metav1.GetOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(dep.GroupVersionKind().String())

		t, err := tableconvertor.TableForObject(dep)
		if err != nil {
			log.Fatalln(err)
		}

		data, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(data))
	}
}
