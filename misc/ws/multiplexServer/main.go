package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/koding/websocketproxy"
	"log"
	"net/http"
	"net/url"
)

var (
	flagBackend = flag.String("backend", "ws://localhost:8080/echo", "Backend URL for proxying")
	upgrader    = websocket.Upgrader{}
)

func main() {
	//u, err := url.Parse(*flagBackend)
	//	//if err != nil {
	//	//	log.Fatalln(err)
	//	//}
	//	//
	//	//http.HandleFunc("/", home)
	//	////http.HandleFunc("/echo", echo)
	//	//
	//	//echoHandler := websocketproxy.ProxyHandler(&url.URL{
	//	//	Scheme: "ws",
	//	//	Path:   "/echo",
	//	//})
	//	//http.Handle("/echo", echoHandler)
	//	//
	//	//log.Println("server running on localhost:8081")
	//	////log.Fatalln(http.ListenAndServe(":8081", websocketproxy.NewProxy(u)))
	//	//log.Fatalln(http.ListenAndServe(":8081", nil))

	endpoint := NewEndpoint()

	endpoint.AddHandleFunc()
}

func home(w http.ResponseWriter, r *http.Request) {

}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Print("read: ", err)
			break
		}

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}
}
