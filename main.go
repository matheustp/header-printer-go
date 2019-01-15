package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(r.Header)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
