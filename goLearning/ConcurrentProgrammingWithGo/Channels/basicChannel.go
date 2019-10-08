package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 10)
	ch <- "Hello"
	str := <-ch
	fmt.Print(str)
}
