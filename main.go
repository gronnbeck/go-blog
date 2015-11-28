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
	Embedded map[string]interface{} `json:"_embedded"`
	Links    map[string]string      `json:"_links"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	embedded := map[string]interface{}{"apple": 5}
	links := map[string]string{"_self": "/posts"}
	hal := HAL{Embedded: embedded, Links: links}
	parsed, _ := json.Marshal(hal)
	fmt.Fprintf(w, string(parsed))
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Envvar PORT missing")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/posts", getPosts)
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	log.Printf("Server running on %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
