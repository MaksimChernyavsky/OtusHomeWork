package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello, OTUS!"
	fmt.Print(reverse.String(str))
}
