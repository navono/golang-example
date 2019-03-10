package composition

import (
	"fmt"
	a "golangTutorialSeries/misc/composition/author"
	p "golangTutorialSeries/misc/composition/post"
	ws "golangTutorialSeries/misc/composition/website"
)

func init() {
	fmt.Println()
	fmt.Println("===> enter composition package")

	// 以下代码可以使用 oop 中的 New 的方式进行构造

	author1 := a.Author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := p.Post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := p.Post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := p.Post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := ws.WebSite{
		Posts: []p.Post{post1, post2, post3},
	}
	w.Contents()

	fmt.Println("<=== exit composition package")
	fmt.Println()
}
