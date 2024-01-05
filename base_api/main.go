package main

import (
	"net/http"
	"os"

	"github.com/robert-min/go-utils/base_api/handlers"
)

func setupServer(mux *http.ServeMux) http.Handler {
	handlers.Register(mux)
	return mux
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8000"
	}

	mux := http.NewServeMux()
	wrappedMux := setupServer(mux)
	http.ListenAndServe(listenAddr, wrappedMux)
}
