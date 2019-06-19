package kubeapi

import (
  // "k8s.io/apimachinery/pkg/api/errors"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/rest"
)

type KubeApiState struct {
  Config *rest.Config
  Clientset *kubernetes.Clientset
}

func New() *KubeApiState {
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err)
  }

  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err)
  }

  return &KubeApiState{
    Config: config,
    Clientset: clientset,
  }
}

func (api KubeApiState) GetNamespaces() []string {
  namespaces, err := api.Clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
  if err != nil {
    panic(err)
  }

  values := []string{}
  for _, item := range namespaces.Items {
    values = append(values, item.Name)
  }

  return values
}

func (api KubeApiState) GetPodCount() int {
  pods, err := api.Clientset.CoreV1().Pods("").List(metav1.ListOptions{})
  if err != nil {
    panic(err)
  }

  return len(pods.Items)
}
