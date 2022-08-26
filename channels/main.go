package main

import (
	"fmt"
	"time"
)

func worker(nThreads int, channel chan string) {
	for i := 0; i <= nThreads; i++ {
		channel <- fmt.Sprintf("Worker %d started\n", i)
		time.Sleep(time.Second)
	}

	close(channel)
}

func main() {
	channel := make(chan string)

	go worker(6, channel)

	fmt.Println("Worker has started creating new threads!")

	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Worker has finished creating new threads!")
}
