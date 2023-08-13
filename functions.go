package main

import (
	"fmt"
	"math/rand"
)

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func greetToPerson(name string) {
	fmt.Println("Hello,", name)
}

func returnsMessage() string {
	return "Sample Message"
}

func addition(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

func returnSingleRandomNumber() int {
	return rand.Intn(100)
}

func returnMultipleRandomNumbers() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func main() {
	dozen := double(6)
	fmt.Println(dozen)
	fmt.Println(add(dozen, 4))
	greet()
	greetToPerson("John")
	fmt.Println(returnsMessage())
	fmt.Println(addition(40, 60, 50))
	fmt.Println(returnSingleRandomNumber())
	fmt.Println(returnMultipleRandomNumbers())
	fmt.Println(double(6) + addition(40, 60, 50)) //162

}
