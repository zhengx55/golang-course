package main

import "fmt"

/*
	1. 基本数据类型, 变量存的就是值
	2. 获取变量的地址 用& var r *int = &r
*/

func main() {
	x := 42
	getAddr(x)
	// fmt.Println(x)
	// fmt.Println(&x)
	// var ptr *int = &x
	// fmt.Println(*ptr)
}

func getAddr(num int) {
	fmt.Printf("num addr = %v\n", &num)
	var ptr *int = &num
	*ptr = 100
	fmt.Printf("num value = %v\n", num)

}

/*

 */
