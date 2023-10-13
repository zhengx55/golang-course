package main

import "fmt"

func main() {
	defer foo()
	bar()
}

// defer 语句用于延迟函数的执行，通常在当前函数返回之前执行被延迟的函数。defer语句允许你确保在函数返回前执行清理操作，例如关闭文件、释放资源等
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
