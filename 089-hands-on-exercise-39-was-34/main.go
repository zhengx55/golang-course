package main

import "fmt"

func main() {
	fmt.Printf("true && true \t\t%v\n", (true && true))   // true
	fmt.Printf("true && false \t\t%v\n", (true && false)) // false
	fmt.Printf("true || true \t\t%v\n", (true || true))   // true
	fmt.Printf("true || false \t\t%v\n", (true || false)) // true
	fmt.Printf("!true \t\t\t%v\n", !true)                 // false
}
