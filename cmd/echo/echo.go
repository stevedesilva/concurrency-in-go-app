package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Timed out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	fmt.Printf(" echo 1")
	l, err := io.Copy(out, in)
	fmt.Printf(" echo 2")
	if err != nil {
		fmt.Errorf("error copying %d bytes: %v", l, err)
	} else {
		fmt.Printf(" copied %d bytes", l)
	}
}
