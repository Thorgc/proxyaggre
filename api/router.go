package api

import (
	"fmt"
	"net/http"
)

const version = "v0.3.10"

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

