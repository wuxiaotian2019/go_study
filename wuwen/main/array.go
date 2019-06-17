package main

import (
	"fmt"
)

func main() {
	//定义一个不设置长度的数组中括号里面32上点[...]
	a := [...]int{5, 2, 3, 1, 4}
	fmt.Println("冒泡排序前数组a=", a)
	//设置变量取数组的长度，在for循环时，每次都要计算数组长度，浪费时间，所以设置一个变量len存储一下
	len := len(a)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if a[i] > a[j] {
				tmep := a[i]
				a[i] = a[j]
				a[j] = tmep
			}
		}
	}
	fmt.Println("冒泡排序后数组a=", a)
}
