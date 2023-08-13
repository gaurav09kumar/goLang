package main

import "fmt"

type Part string

// Create a function to print out the contents of the assembly line
func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func printSlices(title string, slice []string) {
	fmt.Println()
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}

}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlices("Route 1", route)

	route = append(route, "Home")
	printSlices("Route 2", route)

	fmt.Println()
	route[0] = "Visited"
	printSlices("Route 3", route)

	route = route[2:]
	printSlices("Remaining locations", route)

	//Using a slice create an assembly line that contains type Part
	assemblyLine := make([]Part, 3)

	//Create an assembly line having any three parts
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"

	fmt.Println()
	fmt.Println("------------------")
	fmt.Println("3 parts:")
	showLine(assemblyLine)

	fmt.Println("------------------")
	// Add two new parts to the line
	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLine)

	fmt.Println("------------------")
	// Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced:")
	showLine(assemblyLine)

}
