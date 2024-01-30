package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/about", aboutHandler)

	http.Handle("/", r)

	fmt.Println("Server AZEN is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Web!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}


// import (
//     "fmt"
//     "net/http"
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/about", aboutHandler)
// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }