package post

import (
	"fmt"
	"golangTutorialSeries/misc/composition/author"
)

// Post represent a post
type Post struct {
	Title   string
	Content string
	author.Author
}

// Details returns a blog details
func (p Post) Details() {
	fmt.Println("Title: ", p.Title)
	fmt.Println("Content: ", p.Content)
	fmt.Println("Author: ", p.FullName())
	fmt.Println("Bio: ", p.Bio)
}
