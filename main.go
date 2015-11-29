package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gronnbeck/go-blog/hal"

	"strconv"
	"time"
)

type Post struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Ts      time.Time `json:"ts"`
}

func allPosts() []Post {
	posts := make([]Post, 2)
	posts[0] = Post{
		ID:      1,
		Title:   "Hello world",
		Content: "My first post",
		Author:  "Ken Grønnbeck",
		Ts:      time.Now(),
	}
	posts[1] = Post{
		ID:      2,
		Title:   "My first go post",
		Content: "Not done",
		Author:  "Ken Grønnbeck",
		Ts:      time.Now(),
	}

	return posts
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := allPosts()
	links := map[string]interface{}{}
	linkPosts := make([]map[string]string, len(posts))

	for index, post := range posts {
		strID := strconv.Itoa(post.ID)
		linkPosts[index] = map[string]string{
			"href":  "/posts/" + strID,
			"title": post.Title,
		}
	}

	links["posts"] = linkPosts
	embedded := map[string]interface{}{}
	h := hal.HAL{Embedded: embedded, Links: links}
	parsed, _ := json.Marshal(h)
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
