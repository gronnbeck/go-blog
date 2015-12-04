package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gronnbeck/go-blog/comments"
	"github.com/gronnbeck/go-blog/posts"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Envvar PORT missing")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/posts/{postid}/comments", comments.GetCommentsHandler)
	r.HandleFunc("/posts/{id}", posts.GetPostHandler)
	r.HandleFunc("/posts", posts.GetPostsHandler)

	r.HandleFunc("/", handler)

	http.Handle("/", r)

	log.Printf("Server running on %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
