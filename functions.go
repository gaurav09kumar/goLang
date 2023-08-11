package main

import "fmt"

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

func main() {
	dozen := double(6)
	fmt.Println(dozen)
	fmt.Println(add(dozen, 4))
	greet()

	greetToPerson("John")
}
