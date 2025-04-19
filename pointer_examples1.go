package main

import "fmt"

func NoPtr(a int) {
	a = 1
}
func WithPtr(a *int) {
	*a = 1
}

func RunMainPointerExample1() {
	fmt.Println("*******----Start pointers example1 ----***********")

	a := 0
	NoPtr(a)
	println(a)

	WithPtr(&a)
	println(a)
	y := &a
	println(*y)

	fmt.Println("*******----End pointers example1 ----***********")

}
