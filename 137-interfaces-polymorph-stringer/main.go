package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

/*
Sprintf: 用于格式化数据并将其存储为字符串。

Sprintln: 用于格式化数据并添加换行符，然后将其存储为字符串。
*/
func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	// convert int to string
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
