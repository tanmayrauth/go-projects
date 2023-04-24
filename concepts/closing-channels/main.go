package main

import "fmt"

func main() {

	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // more will become true only when the channel is closed
			if more {
				fmt.Println("Job Recieved", j)
			} else {
				fmt.Println("All job has been closed")
				done <- true // essential for blocking main thread to complete
				return
			}
		}

	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Job sent", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done // blocking condition for main

}
