package main

import "fmt"

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- fibonacci(job)
	}
}

func main() {
	jobs := make(chan int, 45)
	results := make(chan int, 45)

	var limit int = 45

	for i := 1; i <= limit; i++ {
		jobs <- i
	}

	close(jobs)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 1; i <= limit; i++ {
		result := <-results
		fmt.Println(result)
	}
}
