package main

import "fmt"

func main() {
	a := 5
	b := &a

	a = 10

	fmt.Println(a, *b)
}
