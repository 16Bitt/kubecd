package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request")
  })

  log.Fatal(http.ListenAndServe(":8000", nil))
}
