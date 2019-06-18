package main

import (
	"fmt"
)

var a string

func swap(a *int, b *int) {
	tem := *a
	*a = *b
	*b = tem
	return
}

func main() {
	a = "G"
	fmt.Println(a)
	f1()
	first := 100
	second := 200
	swap(&first, &second)
	fmt.Println("first = ", first)
	fmt.Println("second = ", second)
}

func f1() {
	a := "O"
	fmt.Println(a)

	f2()
}

func f2() {
	fmt.Println(a)
}
