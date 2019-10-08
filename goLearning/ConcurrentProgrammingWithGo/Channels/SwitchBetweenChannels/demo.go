package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	/*
		msg := Message{
		   		To:      []string{"frodo@underhill.me"},
		   		From:    "gandalf@whitecouncil.org",
		   		Content: "Keep it secret, keep it safe.",
		}

		failedMsg := FailedMessage{
		   		ErrorMessge:     "Message interruptted by black rider",
		   		OriginalMessage: Message{},
		}

		msgCh <- msg
		errCh <- failedMsg
	*/
	//close(msgCh)
	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages at all")
	}
}

//Message template
type Message struct {
	To      []string
	From    string
	Content string
}

//FailedMessage template
type FailedMessage struct {
	ErrorMessge     string
	OriginalMessage Message
}
