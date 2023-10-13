package main

import "fmt"

func main() {
	// 打印一个金字塔
	for i := 1; i <= 3; i++ {
		c := 2*i - 1
		space := 3 - i
		for k := 1; k <= space; k++ {
			fmt.Printf(" ")

		}
		for j := 1; j <= c; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i <= 3; i++ {
		c := 2*i - 1
		space := 3 - i
		for k := 1; k <= space; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= c; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")

			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
