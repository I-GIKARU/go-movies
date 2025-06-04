package main

import (
	"github.com/gorilla/mux"
	"gomysql/pkg/routes"
	"gomysql/pkg/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// Add health check endpoint for Render
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Enable CORS middleware
	handler := utils.EnableCORS(r)

	// Use PORT environment variable, fallback to 9010 for local development
	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
