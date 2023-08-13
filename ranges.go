package main

import "fmt"

//we can go through each word using the range method
// with the ranges we dont have to calculate the length
func main() {
	slice := []string{"Hello", "World", "!"}
	for i, element := range slice {
		fmt.Println(i, element)

		// GO to each letter of each word
		// within this print f we are able to print out a rune instead of just the number of the rune
		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}

}
