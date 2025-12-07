package main

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		// Cara 1
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		// Cara 2
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
