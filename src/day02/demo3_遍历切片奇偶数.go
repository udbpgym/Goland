package main

import "fmt"

type myFunc func(int) bool

func main() {
	arr := []int{23, 4, 5, 6, 7, 9, 10, 345, 789, 124}
	//获取切片中的奇数元素
	odd := Filter(arr, isOdd)
	fmt.Println("奇数元素：", odd)
	even := Filter(arr, isEven)
	fmt.Println("偶数元素：", even)
}

//判断偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

//判断奇数
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

//根据函数来处理切片实现奇数偶数分组返回新的切片
//_代表匿名变量 append相当于在切片后面添加数据相当于add
func Filter(arr []int, f myFunc) []int {
	var resul []int
	for _, value := range arr {
		if f(value) {
			resul = append(resul, value)
		}
	}
	return resul
}
