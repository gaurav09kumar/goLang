package main

import "fmt"

func main() {
	var myName = "john"
	fmt.Println("My Name is", myName)

	var newName string = "kate"
	fmt.Println("My Name is", newName, newName)

	secondName := "henry"
	fmt.Println("Second Name is :", secondName)

	var sum int
	fmt.Println("The sum is", sum) //prints the default value of int

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 4, 9
	fmt.Println("part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum=", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

	//Exercise

	var color = "orange"
	var birthYear, ageInYears = 1998, 25
	var (
		firstInitials  = "A"
		secondInitials = "B"
	)
	var ageInDays int
	ageInDays = 365 * ageInYears

	fmt.Println("Fav color is", color)
	fmt.Println("BirthYear is ", birthYear, "ageInYears is ", ageInYears)
	fmt.Println("FirstInitials is :", firstInitials, "and Second Initials is :", secondInitials)
	fmt.Println("Days old : ", ageInDays)

	//Basic Math operation
	var a, b = 10, 20
	fmt.Println("Multiplication", a*b)
	fmt.Println("Sub", a-b)
	fmt.Println("Add", a+b)

}
