package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %s! Welcome to Go.", *name)
}
