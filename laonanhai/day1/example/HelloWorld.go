package main

import (
	"fmt"
	"strings"
)

func main() {
	//这是我的第一个简单的go程序
	fmt.Println("Hello, World!")
	str := "这里是 www\n.baidu\n.com"
	fmt.Println("---------原字符串-------")
	fmt.Println(str)
	//去除空格
	str = strings.Replace(str, " ", "", -1)
	//去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("---------去除空格和换行后-----------")
	fmt.Println(str)
}
