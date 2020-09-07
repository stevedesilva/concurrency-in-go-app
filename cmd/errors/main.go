package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrTimeout  = errors.New("the request timed out")
	ErrRejected = errors.New("the request was rejected")
	random      = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	response, err := sendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = sendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

}

func sendRequest(req string) (string, error) {
	fmt.Printf("Sending %s\n ", req)
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout

	}
}
