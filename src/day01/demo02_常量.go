package main

import "fmt"

//定义一个结构体
type Man struct {
	age  int
	name string
}

func main() {
	/*
		常量只能是 int string boole
		常量关键字const 格式canst 标识符[类型]=值  类型可以没有
		常量定以后没有被引用，不会在编译时报错 程序运行时不可被修改
		iota 特殊常量值 是一个系统定义的可以被编译器修改的常量值，可以理解为计数器，只要出现常量就开始计数，只要是常量就+1
	*/
	const (
		A = 0
		B = 1
		C = 2
	)
	const (
		a = 10
		b
		c
	)
	fmt.Println(A, B, C, a, b, c)
	const (
		d = iota
		e = iota
		f = iota
	)
	fmt.Println(d, e, f)
	const name = "张三"
	//常量不可被重新复制
	//name="李四"

	man := Man{12, "李四"}
	fmt.Println(man)
	const (
		i = 1 << iota //表示1*2^iota次方  3<<2 =3*2^2
		j = 3 << iota //3*2^1
		k             //3*2^2
		l             //3*2^3
	)
	fmt.Println(i, j, k, l)

	const (
		i1 = '一' //中文字符第一个的ASCII值位19968
		j1
		k1 = iota
		l1
	)
	fmt.Println(i1, j1, k1, l1)

	const sex = iota
	fmt.Println(sex)
}
