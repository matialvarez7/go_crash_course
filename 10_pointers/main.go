package main

import "fmt"

func main() {
	a := 5
	b := &a // save the memory addres of "a"

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read val from addres

	fmt.Println(*b)

	// Change val with pointer

	*b = 10
	fmt.Println(a)
}
