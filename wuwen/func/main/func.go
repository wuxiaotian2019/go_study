package main

import "fmt"
/*
go 没有异常机制，但有panic/recover模式来处理错误
Panic可以在任何地方引发，但是recover 只有在defer调用的函数中有效
*/
func main() {
	A()
	B()
	C()
}
func A()  {
	fmt.Println("Func a")
}

func B()  {
	defer func() {
		if err:=recover();err !=nil{
			fmt.Println("Defer B")
		}
	}()
	panic("Panic in B")
}

func C()  {
	fmt.Println("Func C")
}

/*
因为引用闭包，是匿名函数，闭包是引用地址的
所以在退出defer时，i从0，1，2，执行最后一次i++ 等于3了
output 3 3 3
*/
func main3() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

/*
defer 有个特性，先进后出，后进先出
output a c b
*/
func main2() {

	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func main1() {
	//闭包
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

/*
闭包：定义在一个函数内部的函数，其实就是返回一个匿名函数
*/
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
