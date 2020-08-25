package handlers

import (
	"fmt"
	"net/http"
)

var HelloWorldHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
})
