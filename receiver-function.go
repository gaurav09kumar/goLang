package main

import "fmt"

type Space struct {
	occupied bool
}

//Then we have our parking lot structure, which contains all of the spaces that we have
type ParkingLot struct {
	//we will have our space structure and this is represents a single parking space
	spaces []Space
}

// Regular function
func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

// Create the same function using the receiver function
func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

//Make another receiver function when the vehicle leaves
func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

// Creating a program that can track the status of the parking spaces
func main() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	fmt.Println("initial lot", lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate", lot)
}
