package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	until := time.After(5 * time.Second)
	msg := make(chan string)
	go sender(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Printf("Received %s\n", m)
		case <-until:
			done <- struct{}{}
			return
		}
	}

}

func sender(msg chan<- string, done <-chan struct{}) {
	for {
		select {
		case msg <- "hi":
			time.Sleep(time.Second)
		case <-done:
			fmt.Println("exiting")
			close(msg)
			return
		}
	}
}
