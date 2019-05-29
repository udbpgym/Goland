package main

import "fmt"

/*结构体*/
type point struct {
	x, y int
}

func main() {

	fmt.Println("helloword")
	l := "张三"
	fmt.Printf("%T,%v\n", l, l)
	p := point{1, 2}
	fmt.Printf("%T,%v\n", p, p)
	var a rune = '一'
	fmt.Printf("%T,%v\n", a, a)
	//	布尔值
	m := true
	fmt.Printf("%T,%v\n", m, m)
	//整数
	fmt.Printf("% T,% d \n", 123, 123)
	fmt.Printf("% T,% 5d \n", 123, 123)
	fmt.Printf("% T,% 05d \n", 123, 12345131231)
	//进制转换 二进制
	str := fmt.Sprintf("% b", 123)
	fmt.Print(str)
	//%o 八进制  %U unicode  %x 16进制
	fmt.Printf("% x\n", 123)
	fmt.Printf("% o\n", 123)
	fmt.Printf("% X\n", 123)
	fmt.Printf("% U\n", 123)
	//浮点型 %f==%.6f表示默认保留6位
	fmt.Printf("% f\n", 123.1)
	fmt.Printf("%.2f \n", 123.12222222)
	//科学计数法
	fmt.Printf("%.10e \n", 123.123456)
	//%s直接输出字符串或者字节数组byte (只能是byte数组)
	fmt.Printf("%s \n", "欢迎来到go世界")
	fmt.Printf("%q \n", "欢迎来到go世界")
	arr := [3]byte{97, 98, 99}
	fmt.Printf("%T,%s \n", arr, arr)
	//%x 输出的是16进制索引下标
	fmt.Printf("%T,%x \n", arr, arr)
}
