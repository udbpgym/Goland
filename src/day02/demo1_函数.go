package main

import "fmt"

var a = 7
var b = 8

func main() {
	a, b, c := 10, 20, 0
	fmt.Printf("main函数a %d\n", a)
	fmt.Printf("main函数b %d\n", b)
	fmt.Printf("main函数c %d\n", c)
	//c = sum(a, b)
	sum2(a, b)
}

//两个函数相加
func sum(a, b int) int {
	a++
	b += 2
	c := a + b
	fmt.Printf("sum函数a %d\n", a)
	fmt.Printf("sum函数b %d\n", b)
	fmt.Printf("sum函数c %d\n", c)
	return c
}

//更改返回值起一个变量名
func sum2(a, b int) (c int) {
	a++
	b += 2
	c = a + b
	fmt.Printf("sum函数a %d\n", a)
	fmt.Printf("sum函数b %d\n", b)
	fmt.Printf("sum函数c %d\n", c)
	return
}
