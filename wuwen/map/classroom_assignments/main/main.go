package main

import "fmt"

/*
根据在 for range 部分讲解的知识
尝试将类型为map[int]string的键和值进行交换，变成类型map[string]int
*/
func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println("需要交换的m1 =", m1)
	//定义一个空的map：m2
	m2 := make(map[string]int)
	//使用for range 取出m1的键值对，赋值给m2，做交换
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println("交换后的  m2 =", m2)
}
