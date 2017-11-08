package main

import (
	"fmt"
	"flag"
	"math/rand"
)

func main() {
	// define flags
	maxp :=flag.Int("max", 6, "the max value")
	// parse
	flag.Parse()
	// generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))

	// go run .\commandLine.go max=9
}