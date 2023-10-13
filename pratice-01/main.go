package main

import "fmt"

func main() {
	sum := 0.0
	for i := 0; i < 5; i++ {
		var score float64
		fmt.Printf("请输入第%d个学生的成绩\n", i)
		fmt.Scanln(&score)
		sum += score
	}
	fmt.Printf("Average %v\n", sum/5)
}
