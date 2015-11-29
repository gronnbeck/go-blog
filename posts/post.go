package posts

import "time"

// Post is the datastructure for exposing posts through the API
type Post struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Ts      time.Time `json:"ts"`
}
