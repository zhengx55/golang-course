package main

import "fmt"

func main() {
	// si := []string{"a", "b", "c"}
	// fmt.Println(si)

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println("------------")
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------")
	mySlice := []int{1, 2, 3}
	fmt.Printf("切片的容量: %d\n", cap(mySlice))

	// 向切片添加元素
	mySlice = append(mySlice, 4)
	fmt.Printf("切片的容量: %d\n", cap(mySlice))
	fmt.Printf("切片的长度: %d\n", len(mySlice))

}
