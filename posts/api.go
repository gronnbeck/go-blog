package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	parsed, _ := json.Marshal(h)
	fmt.Fprintf(w, string(parsed))
}
