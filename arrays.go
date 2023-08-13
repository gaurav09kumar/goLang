package main

import "fmt"

type Room struct {
	name    string
	cleaned bool //false is the default values for boolean
}

//So the array that we use has to have four elements and it must be of type room
func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i] //** Alaways make sure whenever you utilize the iterator variable within a loop, you are making it in the context of a copy
		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

type Product struct {
	name  string
	price int
}

func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}

	}
	fmt.Println("Last item on the list:", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost of each items", cost)
}

func main() {

	//Make Array for Rooms
	rooms := [...]Room{
		{name: "Office"},
		{name: "WareHouse"},
		{name: "Reception"},
		{name: "Operations"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning")
	rooms[2].cleaned = true //3rd room is cleaned
	rooms[3].cleaned = true //4th room is cleaned

	checkCleanliness(rooms)

	fmt.Println("+++++++++++++++++++EXERCISE+++++++++++++++++++++")

	shoppingList := [4]Product{
		{"Rice", 100},
		{"Dal", 50},
	}

	printStats(shoppingList)

	shoppingList[2] = Product{"Bread", 5}

	printStats(shoppingList)

}
