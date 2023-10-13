package main

import (
	"io"
	"log"
	"os"
)

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

type person struct {
	first_name string
	last_name  string
	age        int
	gender     string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first_name))
	return err
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).

		log.Fatalf("error %s", err)
	}
	defer f.Close()
	s := []byte("Hello gophers!")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
