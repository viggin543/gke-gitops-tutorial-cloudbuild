package main

import (
	"github.com/viggin543/gke-gitops-tutorial-cloudbuild/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/banana", handlers.HelloServer)
	_ = http.ListenAndServe(":8080", nil)
}

