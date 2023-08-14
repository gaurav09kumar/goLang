package main

import (
	"fmt"
)

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

const (
	Active   = true
	InActive = false
)

//Create type alias , will help us manage our code shortly
type SecurityTag bool

//Create a structure to store items amd their security tag state
type Item struct {
	name string
	tag  SecurityTag
}

// Security tags have two states: active (true) and inactive(false)
// Create a functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = InActive
}

func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

//Create a checkout
// In this demo will be creating two functions that utilize pointers and will be changing data at different points in the program
func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

	fmt.Println("++++++++++EXERCISE+++++++++++++")
	//Security system that can deactivate and activate tags on products
	//Create at least 4 items, all with active security tags
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	watch := Item{"watch", Active}

	// - Store them in a slice or array
	items := []Item{shirt, pants, purse, watch}
	fmt.Println(items)

	// Deactivate any one of the security tags within the array / slice
	deactivate(&items[0].tag) // creating a pointer to the security item - shirt which will be deactivated
	fmt.Println("Items 0 deactivated", items)

	//Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println("checked out", items)
}
