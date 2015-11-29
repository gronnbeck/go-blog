package posts

import "time"

func getAllPosts() []Post {
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
