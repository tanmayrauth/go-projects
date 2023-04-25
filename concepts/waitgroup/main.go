package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("Starting worker %d \n", i)
	time.Sleep(time.Second)
	fmt.Printf("Ending worker %d \n", i)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // This increments value in wg
		// i := i
		go func() {
			defer wg.Done() // This decrements value in wg
			worker(i)
		}()

	}

	wg.Wait() // This holds the process till the wg value is not 0
}
