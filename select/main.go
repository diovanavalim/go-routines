package main

import (
	"fmt"
	"time"
)

func worker(channel chan string, threadId int, qt time.Duration) {
	for {
		time.Sleep(qt)
		channel <- fmt.Sprintf("Channel %d", threadId)
	}
}

func main() {
	channel1 := make(chan string)

	channel2 := make(chan string)

	go worker(channel1, 1, time.Millisecond*500)

	go worker(channel2, 2, time.Second*2)

	for {
		select {
		case msg1 := <-channel1:
			fmt.Println(msg1)
		case msg2 := <-channel2:
			fmt.Println(msg2)
		}

	}
}
