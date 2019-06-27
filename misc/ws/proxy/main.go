package main

import (
	"flag"
	"github.com/koding/websocketproxy"
	"log"
	"net/http"
	"net/url"
)

var (
	flagBackend = flag.String("backend", "ws://localhost:8080/echo", "Backend URL for proxying")
)

func main() {
	u, err := url.Parse(*flagBackend)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("proxy server running on localhost:9090")
	log.Fatalln(http.ListenAndServe(":9090", websocketproxy.NewProxy(u)))
}
