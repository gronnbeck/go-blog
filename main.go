package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"_embedded\": { \"posts\": [] }}")
}

func main() {
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	r.HandleFunc("/posts", getPosts)
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	log.Printf("Server running on %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
