package main

import "fmt"

func main() {
	/*
		与java类似只是多了一个select语句
	*/
	EvenOdd()
	avgNum(70)
}

//首字母大写表示公有
func EvenOdd() {
	//if条件语句
	a := 10
	b := 11
	if a == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")

	}
	if a%2 == 0 {
		fmt.Println(a, "是偶数")
	} else {
		fmt.Println(a, "是奇数")
	}
}

//首字母小写表示私有
func avgNum(souce int) {

	if souce >= 90 {
		fmt.Println("优秀")
	} else if 90 > souce && souce >= 80 {
		fmt.Println("优良")
	} else if 80 > souce && souce >= 70 {
		fmt.Println("中等")
	} else if 70 > souce && souce >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
