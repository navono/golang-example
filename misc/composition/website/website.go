package website

import (
	"fmt"
	p "golangTutorialSeries/misc/composition/post"
)

// WebSite represent a web site
type WebSite struct {
	Posts []p.Post
}

// Contents returns all post of web site
func (w WebSite) Contents() {
	fmt.Println("Contents of WebSite")
	for _, v := range w.Posts {
		v.Details()
		fmt.Println()
	}
}
