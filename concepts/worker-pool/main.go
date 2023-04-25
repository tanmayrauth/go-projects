package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("Worker started for :", j)
		time.Sleep(time.Second * 3)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numJobs; w++ {
		go worker(jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for r := 1; r <= numJobs; r++ {
		<-results
	}

}
