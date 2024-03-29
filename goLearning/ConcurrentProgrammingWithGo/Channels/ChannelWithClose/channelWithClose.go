package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "There are the times that try men's soul.\n"

	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word

	}

	close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Printf(<-ch + " ")
	}

	ch <- "Test efter ch is closed"

}
