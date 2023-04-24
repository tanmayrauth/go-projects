package main

import (
	"fmt"
)

type Rectangle struct {
	height int
	width  int
}

func (r *Rectangle) calculateArea() int {
	return r.height * r.width
}

func DoubleHeight(rect *Rectangle) {
	rect.height = rect.height * 2
}

func apple() func() int {
	i := 2
	return func() int {
		return i + 1
	}
}

func main() {
	fmt.Println("Hello World")

	msg := make(chan string)

	go func() { msg <- "apple" }()

	val := <-msg

	r := Rectangle{
		height: 3,
		width:  4,
	}
	// ptr := &r
	fmt.Println(r.calculateArea())
	DoubleHeight(&r)
	fmt.Println(r.height)

	fmt.Println(val)

	// fval := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("This is boolean")

	// 	case int:
	// 		fmt.Println("This is int")
	// 	default:
	// 		fmt.Printf("Type is %T\n", t)
	// 	}
	// }
	// fval(1)

	fmt.Printf("Apple value type is %T", apple())

}
