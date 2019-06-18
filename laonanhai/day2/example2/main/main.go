package main

import (
	"day2/example2/add"
	"fmt"
)

func init() {
	fmt.Println("this is main.init")
}

func main() {
	fmt.Println("main add.Name =", add.Name)
	fmt.Println("main")
}
