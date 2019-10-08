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
	/* 	
	for msg := range ch {
		fmt.Printf(msg + " ")
	} 
	*/
	for {
		if msg, ok := <-ch; ok {
			fmt.Printf(msg + " ")
		} else {
			break
		}
	} 


}
