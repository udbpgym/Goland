package main

import "fmt"

func main() {
	Mark("B")
}

/*
GO语言中switch会给每一个语句自带break，如果想贯穿可以添加 fallthrough
case 可以添加多个条件，所有case不能出现相同的元素，用逗号分割
*/
func Mark(souce string) {
	switch souce {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
		//fallthrough
	case "C":
		fmt.Println("及格")
		//fallthrough
	case "D":
		fmt.Println("不及格")

	default:
		fmt.Println("差")
	}
}
