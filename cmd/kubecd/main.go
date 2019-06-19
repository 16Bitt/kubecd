package main

import (
  "fmt"
  "log"
  "net/http"
  "app/internal/kubeapi"
)

func main() {
  api := kubeapi.New()
  count := api.GetPodCount()
  fmtCount := fmt.Sprintf("%d pods", count)

  http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(fmtCount))
  })

  log.Fatal(http.ListenAndServe(":8000", nil))
}
