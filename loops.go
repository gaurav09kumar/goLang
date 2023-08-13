package main

import "fmt"

func fizzBuzz() {
	i := 1
	for i <= 50 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
		i++
	}
}

func main() {
	sum := 0
	fmt.Println("Sum is", sum)

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

	//Exercises
	fizzBuzz()

	fmt.Println("**********************************")
	for j := 1; j <= 50; j++ {
		divisibleBy3 := j%3 == 0
		divisibleBy5 := j%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println("FizzBuzz")
		} else if divisibleBy3 {
			fmt.Println("Fizz")
		} else if divisibleBy5 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}
