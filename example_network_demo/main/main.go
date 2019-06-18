package main

import (
	"fmt"
	"time"
)

func f() chan interface{} {
	c := make(chan interface{})
	go func() {
		fmt.Println(<-c)
	}()
	return c
}

type S struct {
	C chan interface{}
}

func main() {
	a := S{f()}
	a.C <- "哈哈哈"
	time.Sleep(2 * time.Second)
	//output 哈哈哈
}
