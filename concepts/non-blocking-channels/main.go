package main

import "fmt"

func main() {

	msg := make(chan string)
	sig := make(chan bool)

	// This is non blocking receive
	select {
	case m := <-msg:
		fmt.Println("msg received", m)
	default:
		fmt.Println("No Msg received")
	}

	// This is non blocking send
	val := "New Msg"
	select {
	case msg <- val:
		fmt.Println("Message sent")
	default:
		fmt.Println("Message failed to send")

	}

	// Multiple non blocking receive
	select {
	case v := <-msg:
		fmt.Println("Non blocking receive", v)
	case s := <-sig:
		fmt.Println("Signal received", s)
	default:
		fmt.Println("No Activity")

	}
}
