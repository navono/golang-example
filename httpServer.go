package main

import "net/http"
import "io"

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello</title>
			</head>
			<body>
				Hello, Yo!
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	// http.Handle()
	// http.StripPrefix()

	http.ListenAndServe(":9000", nil)
}
