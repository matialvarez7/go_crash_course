package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids

	for i, id := range ids { // i represent the index and id the value
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids { // when we are not using the index use the "_"
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range whith maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// only the key
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
