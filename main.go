package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
)

type HAL struct {
	Embedded string `json:"_embedded"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	hal := HAL{"hello"}
	parsed, _ := json.Marshal(hal)
	fmt.Fprintf(w, string(parsed))
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
