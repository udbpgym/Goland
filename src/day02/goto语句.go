package main

import "fmt"

/*
能用break以及continu时尽量不适用goto
*/
func main() {
	eludeFourGoto()
}
func eludeFourGoto() {
	num := 0
Loop:
	for num < 50 {
		num++
		if num%10 == 4 {
			//满足条件跳转到loop
			goto Loop
		}
		fmt.Printf("%d\t", num)
	}
}
