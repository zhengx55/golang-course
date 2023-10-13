package main

import "fmt"

func main() {
	i := 20
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}

func test() {
	i := 20
	for i >= 0 {
		fmt.Println((i))
		i--
	}
}
