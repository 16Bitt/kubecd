package main

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "app/internal/kubeapi"
)

func main() {
  api := kubeapi.New()

  http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
    count := api.GetPodCount()
    fmtCount := fmt.Sprintf("%d pods", count)

    w.Write([]byte(fmtCount))
  })

  http.HandleFunc("/namespaces", func(w http.ResponseWriter, r *http.Request) {
    namespaces := api.GetNamespaces()
    joined := strings.Join(namespaces, "\n")
    w.Write([]byte(joined))
  })

  log.Fatal(http.ListenAndServe(":8000", nil))
}
