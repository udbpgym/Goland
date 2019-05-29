package main

import "fmt"

func main() {
	/*
		与java类似，不过结束符多了一个goto，for条件绝对不允许出现括号
		结构 for 初始init:条件表达式condition;结束语句post{
					循环体
		}
		for条件三个组成部分为可选
	*/
	//For1()
	//For2()
	//For3()
	//For4()
	For5()

	str := "helloWord"
	for i, char := range str {
		fmt.Println(i, char)
	}
}

//常见用法
func For1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}

//省略初始语句
func For2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i)
	}
}

//省略条件表达式
func For3() {
	i := 0
	for ; ; i++ {
		if i > 10 {
			break
		}
		fmt.Print(i)
	}
}

//省略pst
func For4() {

	for i := 0; i < 20; {
		i++
		fmt.Print(i)
	}
}
func For5() {
	str := "123abcABC"
	for i, value := range str {
		fmt.Printf("第%d位的字符是：%v ,字符是%c \n", i, value, value)
	}
}
