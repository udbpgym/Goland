package main

import (
	"fmt"
	"strings"
)

/*
字符串奇偶交替
*/

//使用type关键词来代替函数名
type caseFunc func(string) string

func main() {
	str := "abcdEfga"
	//result := processLetter(str)
	//fmt.Println(result)
	//使用函数作为参数
	strs := StringToCase(str, processLetter)
	//使用type代替函数作为参数
	strs2 := StringToCase2(str, processLetter)
	fmt.Println(strs)
	fmt.Println(strs2)
}

//函数名小写字母表示私有
func processLetter(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			//strings.ToUpper大写
			result += strings.ToUpper(string(value))
		} else {
			//strings.ToLower 小写
			result += strings.ToLower(string(value))
		}
	}
	return result
}

//函数名大写字母表示公有   把函数当一个变量来使用  myfunc 参数名 func(string) string参数类型
func StringToCase(str string, myfunc func(string) string) string {
	return myfunc(str)
}
func StringToCase2(str string, myfunc caseFunc) string {
	return myfunc(str)
}
