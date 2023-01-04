package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArray [2]string

	// // Assing values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	//Declare and assign shorthand
	fruitArray := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArray)

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice)) // longitud de un slice y array

}
