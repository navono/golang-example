package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var (
	addr = flag.String("addr", "localhost:9090", "http service address")
)

func main() {
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	//d := websocket.Dialer{
	//	Proxy: internalProxy,
	//	//Proxy: http.ProxyURL(&url.URL{
	//	//	Scheme: "http", // or "https" depending on your proxy
	//	//	Host:   "localhost:9090",
	//	//	Path:   "/echo",
	//	//}),
	//	HandshakeTimeout: 45 * time.Second,
	//}
	//ws, _, err := d.Dial(u.String(), nil)

	//backend, err := url.Parse("localhost:9090")
	//NewProxy(backend)

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := ws.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func internalProxy(r *http.Request) (*url.URL, error) {
	log.Println("proxy: ", r.URL)

	//return nil, nil
	return http.ProxyURL(&url.URL{
		Scheme: "http", // or "https" depending on your proxy
		Host:   "localhost:9090",
		Path:   "/echo",
	})(r)
}
