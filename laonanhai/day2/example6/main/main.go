package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

/*
获取当天时间能整除2，打印female，不能打印man
*/
func main() {

	for {
		Second := time.Now().Unix()
		if Second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(time.Second)
	}
}
