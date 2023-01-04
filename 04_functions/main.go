package main

import "fmt"

func greeting(name string) string { //declaramos la función con el valor de retorno que tiene
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Matias"))
	fmt.Println(getSum(5, 8))
}
