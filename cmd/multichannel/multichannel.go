package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

const (
	timeoutDuration = 5 * time.Second
	success         = 0
)

func main() {
	// read from input and write to output using chan
	// could use done := time.After(timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	// timeout after n secs
	echo := make(chan []byte)

	// listen for input from stdin
	go stdInput(echo)
	// write output to fmt.Println or timeout
	for {
		select {
		case val := <-echo:
			fmt.Printf("Writing %v \n", string(val))
			_, err := os.Stdout.Write(val)
			if err != nil {
				fmt.Printf("Unable to write bytes %s\n", err)
				continue
			}
		case t := <-ctx.Done():
			fmt.Printf("Timeout %v \n", t)
			os.Exit(success)

		}
	}
}

func stdInput(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, err := os.Stdin.Read(data)
		if err != nil {
			fmt.Printf("Unable to read bytes %s\n", err)
			continue
		}
		if l > 0 {
			out <- data
		}
	}
}
