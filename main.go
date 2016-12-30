package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./hello.so")
	if err != nil {
		log.Fatal(err)
	}

	ps, err := p.Lookup("Hello")
	if err != nil {
		log.Fatal(err)
	}

	hello, ok := ps.(func(string) string)
	if !ok {
		log.Println("malformed plugin")
	}

	fmt.Println(hello("Alex"))
}
