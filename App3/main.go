// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Item struct represents an item with a name and price.
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

// GetItemsHandler returns the list of items.
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// AddItemHandler adds a new item to the list.
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newItem.ID = fmt.Sprintf("%d", len(items)+1)
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	router := mux.NewRouter()

	// CORS middleware (for development, allow all origins)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.Use(handlers.CORS(headers, methods, origins))

	// API routes
	router.HandleFunc("/api/items", GetItemsHandler).Methods("GET")
	router.HandleFunc("/api/items", AddItemHandler).Methods("POST")

	// Serve static files (React build)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/build")))

	fmt.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
