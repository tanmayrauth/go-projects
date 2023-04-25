package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case d := <-t.C:
				fmt.Println("Ticker at ", d)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	t.Stop()
	done <- true
	fmt.Println("Ticker Stopped")

}
