package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gronnbeck/go-blog/hal"
)

// GetPostsHandler handles API calls for exposing posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts := getAllPosts()
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
	parsed := hal.JSON(h)
	fmt.Fprintf(w, parsed)
}

// GetPostHandler exposes the post fetched from url var
func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Fprintf(w, "Id %v is not an integer", idStr)
		return
	}

	post, err := fetchPost(id)

	if err != nil {
		fmt.Fprintf(w, "The post with id %v does not exist", idStr)
		return
	}

	links := map[string]interface{}{}
	links["self"] = "/posts/" + idStr
	links["comments"] = "/posts/" + idStr + "/comments"

	embedded := map[string]interface{}{}
	embedded["comments"] = make([]map[string]interface{}, 0)

	h := hal.HAL{Embedded: embedded, Links: links, Data: *post}
	parsed := hal.JSON(h)
	fmt.Fprintf(w, parsed)
}
