package main

import "fmt"

func main() {
	// Using var
	var name string = "Matias"
	var age int = 31
	var isCool = true

	// Shortcut
	name2 := "Gabriel" // esta forma solo sirve para dentro de un bloque de función
	fmt.Println(name, age, isCool, name2)
	fmt.Printf("%T\n", age)
}
