package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

var users = []User{
	{Name: "Miguel Tineo", Role: "Senior Engineering Manager"},
	{Name: "Marta Vadilonga", Role: "UI/UX expert"},
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]User{"users": users})
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/records", getRecords).Methods("GET")

	// Use the CORS middleware
	corsRouter := enableCORS(router)

	log.Println("Server running on port 3001")
	log.Fatal(http.ListenAndServe(":3001", corsRouter))
}
