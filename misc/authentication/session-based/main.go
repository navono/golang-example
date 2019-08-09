package session_auth

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var cache redis.Conn

func init() {
	initCache()

	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}

	cache = conn
}
