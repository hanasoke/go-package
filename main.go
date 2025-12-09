package main

import (
	"encoding/json"
	"go-package/config"
	"go-package/controller"
	"go-package/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	// Verifikasi koneksi database
	db := config.DB
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.Index).Methods("GET")

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(":"+config.ENV.PORT, r)
}
