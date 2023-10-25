package main

import (
	"log"
	"net/http"
	"tes-mnc/configs"
	"tes-mnc/routes"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
