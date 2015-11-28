package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"

	"strconv"
	"time"
)

type HAL struct {
	Embedded map[string]interface{} `json:"_embedded"`
	Links    map[string]interface{} `json:"_links"`
}

type Post struct {
	Id      int       `json:id`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Ts      time.Time `json:"ts"`
}

func allPosts() []Post {
	posts := make([]Post, 2)
	posts[0] = Post{
		Id:      1,
		Title:   "Hello world",
		Content: "My first post",
		Author:  "Ken Grønnbeck",
		Ts:      time.Now(),
	}
	posts[1] = Post{
		Id:      2,
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
		strID := strconv.Itoa(post.Id)
		linkPosts[index] = map[string]string{
			"href":  "/posts/" + strID,
			"title": post.Title,
		}
	}

	links["posts"] = linkPosts
	embedded := map[string]interface{}{}
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
