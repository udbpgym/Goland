package main

import (
	"fmt"
	"math"
)

func main() {
	//1 定义调用无参匿名函数
	func(num int) {
		fmt.Println("打印数字是：", num)
	}(100)

	// 定义调用有参匿名函数  math.Sqrt求取平方根
	result := func(num float64) float64 {
		return math.Sqrt(num)
	}(98)
	fmt.Println("平方根：", result)

	//将匿名函数赋值给变量，需要时在调用
	myfunc := func(data string) string {
		return data
	}
	fmt.Println(myfunc("欢迎学习GO语言"))
}
