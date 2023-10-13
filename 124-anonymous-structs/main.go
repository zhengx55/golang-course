package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}
	p2.age = 22

	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
	fmt.Printf("%v \n", p2)
}
