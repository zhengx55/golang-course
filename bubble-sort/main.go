package main

import "fmt"

func main() {
	arr := [5]int{22, 53, 11, 27, 55}
	bubbleSort(&arr)
}

func bubbleSort(arr *[5]int) {
	fmt.Println(*arr)
	temp := 0
	fmt.Println(len(*arr))
	for range *arr {
		for i := 0; i < len(*arr)-1; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}

	}
	fmt.Println(*arr)

}
