package main

import "fmt"

func main() {
	chinese := 90
	english := 80.5

	/*
		平均值avg 两个变量类型必须一样才能进行计算
		如果想计算必须变量进行转换，采用T（表达式）
		会存在精度丢失float转换int精度丢失
		布尔值无法跟任何基本类型转换
	*/
	//float可以用int转换
	avg := (chinese + int(english)) / 2
	//int 转换float 选择float32 或者float64必须根据电脑位数。64位必须用float64
	avg2 := (float64(chinese) + english) / 2
	fmt.Println(avg)
	fmt.Println(avg2)
	//string 不能转换int
	//str := "张三"
	//int(str)
	//int转换string 会直接转换成十进制
	result := string(chinese)
	fmt.Println(result)

}
