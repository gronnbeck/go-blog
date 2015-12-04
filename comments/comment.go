package comments

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gronnbeck/go-blog/hal"
)

// Comment is a generic datastructure for comments
type Comment struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Ts      time.Time `json:"ts"`
}

func (c Comment) toHAL() hal.HAL {
	links := map[string]interface{}{}
	embedded := map[string]interface{}{}
	h := hal.HAL{Embedded: embedded, Links: links, Data: c}
	return h
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["postid"]
	id, err := strconv.Atoi(idstr)

	if err != nil {
		fmt.Fprintf(w, "Id %v is not an integer", idstr)
		return
	}

	comments := GetComments(id)
	links := map[string]interface{}{}
	embedded := map[string]interface{}{}

	embedded["comments"] = translateHals(halifyComments(comments))
	h := hal.HAL{Embedded: embedded, Links: links}
	j := hal.JSON(h)
	fmt.Fprintf(w, j)
}

func halifyComments(comments []Comment) []hal.HAL {
	hals := make([]hal.HAL, len(comments))
	for i, c := range comments {
		h := c.toHAL()
		hals[i] = h
	}
	return hals
}

func translateHals(hals []hal.HAL) []map[string]interface{} {
	translated := make([]map[string]interface{}, len(hals))
	for i, h := range hals {
		translated[i] = hal.Translate(h)
	}
	return translated
}

func GetComments(postId int) []Comment {
	comments := make([]Comment, 1)
	comments[0] = Comment{
		ID:      postId,
		Content: "This is a comment",
		Author:  "Ken Grønnbeck",
		Ts:      time.Now(),
	}
	return comments
}
