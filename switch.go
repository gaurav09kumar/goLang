package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Costly item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First class seating")
	default:
		fmt.Println("Other seating")
	}

	//Exercise

	switch age := 10; {
	case age == 0:
		fmt.Println("newborn")
	case age >= 1 && age <= 3:
		fmt.Println("toddler")
	case age >= 4 && age <= 12:
		fmt.Println("child")
	case age >= 13 && age <= 17:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
