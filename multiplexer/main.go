package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplexer(worker("Hello, World!"), worker("Nice to meet you!"))

	for i := 0; i <= 10; i++ {
		fmt.Println(<-channel)
	}
}

func worker(txt string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Text: %s", txt)
			time.Sleep(time.Second)
		}
	}()

	return channel
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		for {
			select {
			case msg := <-channel1:
				output <- msg
			case msg := <-channel2:
				output <- msg
			}
		}
	}()

	return output
}
