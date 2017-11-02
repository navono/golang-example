package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type String string
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, str)
}

func (str Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "%v %v %v\n", str.Greeting, str.Punct, str.Who)
}

func main() {
	// var h Hello
	// err := http.ListenAndServe("localhost:4000", h)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.Handle("/string", String("I'm yours"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
