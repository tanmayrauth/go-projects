package main

import "fmt"

func main() {

	buffer := make(chan int, 2)
	buffer <- 1
	buffer <- 3
	close(buffer) // not closing the channel lead to deadlock

	for j := range buffer { // here we're treating a channel just like a normal datastructure like list, slice
		fmt.Println(j)
	}
}
