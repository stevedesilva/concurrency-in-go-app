package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Printf("defer will be called at the end regardless what happens\n")
	}()
	callBadFunc()
	fmt.Printf("End Main\n")
}

func callBadFunc() {
	var msg = "Steve"
	defer func() {
		fmt.Printf("<d>callBadFunc %s\n", msg)
	}()

	defer func() {
		r := recover()
		fmt.Printf("<d>recover: %v\n", r)
	}()
	panicFunc()
	fmt.Printf("callBadFunc: Won't be called\n")

}

func panicFunc() {
	defer func() {
		fmt.Printf("panicFunc defer will be called at the end regardless what happens\n")
	}()
	panic(errors.New("something went wrong"))
	fmt.Printf("panicFunc: Won't be called\n")
}
