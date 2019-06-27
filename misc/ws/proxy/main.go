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
	err = http.ListenAndServe(":9090", websocketproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}

	//listener, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		//handle error.
	//		log.Fatalf("failed listening for '%v' on %v: %v", "localhost", 8080, err)
	//	}
	//	go handleConn(conn)
	//}
}

//
//func handleConn(from net.Conn) {
//
//}
