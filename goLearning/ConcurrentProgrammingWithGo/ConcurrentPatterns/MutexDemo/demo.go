package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 100)
	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + "-" + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mutex <- true
			var a int = i
			var b int = j
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", a, b, a+b)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}

	//fmt.Scanln()

	//use channel to replace mutex
	/*
		runtime.GOMAXPROCS(4)
		mutex := make(chan bool, 1)

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				mutex <- true
				var a int = i
				var b int = j
				go func() {
					fmt.Printf("%d + %d = %d\n", a, b, a+b)
					<-mutex
				}()
			}
		}

		fmt.Scanln()
	*/

	/*
		//Use syn.mutex as sync lock
		runtime.GOMAXPROCS(4)
		mutex := new(sync.Mutex)
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				mutex.Lock()
				var a int = i
				var b int = j
				go func() {
					fmt.Printf("%d + %d = %d\n", a, b, a+b)
					mutex.Unlock()
				}()
			}
		}

		fmt.Scanln()
	*/

}
