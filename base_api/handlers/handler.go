package handlers

import (
	"fmt"
	"net/http"
)

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Hello World!!")
}
