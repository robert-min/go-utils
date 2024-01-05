package handlers

import "net/http"

// Register set HandlerFunc
func Register(mux *http.ServeMux) {
	mux.HandleFunc("/hello", getHelloWorld)
	mux.HandleFunc("/name", postNameHandler)
}
