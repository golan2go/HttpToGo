package main

import (
	"encoding/json"
	"fmt"
	"git.att.com/nlp"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	if err := checkHealth(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/_/health", healthHandler)
	http.HandleFunc("/_/tokenize", tokenizeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not GET", http.StatusMethodNotAllowed)
	}
	if err := checkHealth(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "OK\n")
}

func checkHealth() error {
	//check check check
	return nil
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tokens := nlp.Tokenize(string(data))

	resp := map[string]interface{}{
		"error":  nil,
		"tokens": tokens,
	}

	json.NewEncoder(w).Encode(resp)

}
