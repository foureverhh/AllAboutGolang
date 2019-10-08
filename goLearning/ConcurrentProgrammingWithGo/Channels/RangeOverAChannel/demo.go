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

	close(ch) //Only close the sender!!

	/* 	for {
		if msg, ok := <-ch; ok {
			fmt.Printf(msg + " ")
		} else {
			break
		}
	} */

	for msg := range ch {
		fmt.Printf(msg + " ")
	}
}
