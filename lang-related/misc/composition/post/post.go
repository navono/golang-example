package post

import (
	"fmt"
	author2 "golang-example/lang-related/misc/composition/author"
)

// Post represent a post
type Post struct {
	Title   string
	Content string
	author2.Author
}

// Details returns a blog details
func (p Post) Details() {
	fmt.Println("Title: ", p.Title)
	fmt.Println("Content: ", p.Content)
	fmt.Println("Author: ", p.FullName())
	fmt.Println("Bio: ", p.Bio)
}
