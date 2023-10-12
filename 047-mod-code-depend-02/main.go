package main

import (
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)

	s5 := dog.WhenGrownUp("wu")
	fmt.Println((s5))

	// also like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
