package main

import (
	"fmt"
	"time"
)

const maxJob = 5

func main() {
	done := make(chan struct{}, 2)
	lockWithoutTimeout(done)
	withTimeout(done)
	for i := 0; i < 2; i++ {
		<-done
	}
}

func withTimeout(d chan struct{}) {
	// ch
	ch := make(chan bool, 1)
	// start go routines
	for i := 1; i <= maxJob; i++ {
		go work(ch, i)
	}
	// wait then exit
	time.Sleep(5 * time.Second)
	fmt.Println("End withTimeout")
	d <- struct{}{}
}

func work(lock chan bool, num int) {
	fmt.Printf("Attempt get lock for %d\n", num)
	// do some work
	lock <- true
	fmt.Printf("%d has the lock\n", num)
	fmt.Printf("%d working ...\n", num)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Releasing lock for %d\n", num)
	// release lock
	<-lock
}

func lockWithoutTimeout(d chan struct{}) {
	// ch
	ch := make(chan bool, 1)
	done := make(chan bool)
	// start go routines
	for j := 1; j <= maxJob; j++ {
		go workDoneCh(ch, done, j)
	}

	// wait then exit
	for i := 1; i <= maxJob; i++ {
		<-done
		fmt.Printf("done %d \n", i)

	}

	fmt.Println("End lockWithoutTimeout")
	d <- struct{}{}
}

func workDoneCh(lock, done chan bool, num int) {
	fmt.Printf("Attempt get lock for %d\n", num)
	// do some work
	lock <- true
	fmt.Printf("%d has the lock\n", num)
	fmt.Printf("%d working ...\n", num)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Releasing lock for %d\n", num)
	// release lock
	<-lock
	done <- true
}
