package main

import "fmt"

// Arrays are mostly used as a building block in the Go programming languange. In some instances we might use an array, but mostly the recommendation is to use slices instead
func main() {

	// array literal
	a := [3]int{42, 43, 44}
	// automatically store 0 in uninitialized position
	a2 := [4]int{2, 2, 2}
	fmt.Println(a)
	fmt.Println(a2)

	b := [...]string{"Hello", "Gophers!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	// c = a

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42

		// %#v 是Go语言中用于格式化输出的格式化字符串，它通常用于打印Go值的Go语法表示形式，包括类型信息。这可以帮助你更详细地了解一个值的类型和结构。
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
}

/*
"Arrays have their place, but they’re a bit inflexible,
so you don’t see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide
great power and convenience."
~ Go Slices: usage and internals - Go Blog - Andrew Gerrand
*/
// https://go.dev/blog/slices-intro

// PRACTICE - next video
/*
Use a composite literal array
to store these elements in an array
and let the compiler determine the length of the array
*/
