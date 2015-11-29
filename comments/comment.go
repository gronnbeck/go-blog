package comments

import "time"

// Comment is a generic datastructure for comments
type Comment struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
	Ts      time.Time `json:"ts"`
}
