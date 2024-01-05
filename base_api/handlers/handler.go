package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	Name   string `json: "name"`
	Number int    `json: "number"`
}

type Response struct {
	Result string `json: "result"`
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Hello World!!")
}

func postNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var userData UserData
	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&userData)
	if err != nil {
		http.Error(w, "Invalid Json payload.", http.StatusBadRequest)
	}

	resultNumber := userData.Number * 2
	response := Response{
		Result: fmt.Sprintf("%s %d", userData.Name, resultNumber),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
