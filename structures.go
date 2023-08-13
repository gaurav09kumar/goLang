package main

import "fmt"

//Passenger struct
type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

//Bus struct has Passenger
type Bus struct {
	FrontSeat Passenger
}

//Rectangle
type Rectangle struct {
	Length float32
	Width  float32
}

func calculateRectangleAreaPerimeter(rectangle Rectangle) (float32, float32) {
	return rectangle.Length * rectangle.Width, 2 * (rectangle.Length + rectangle.Width)
}

func main() {
	john := Passenger{"john", 123, false}
	fmt.Println(john.Name, john.TicketNumber, john.Boarded)

	//create few structures using the var keyword
	var (
		bill = Passenger{Name: "Bill", TicketNumber: 505}
		ella = Passenger{Name: "Ella", TicketNumber: 775}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	john.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has Boarded>>!")
	}
	if john.Boarded {
		fmt.Println("John has Boarded>>!")
	}

	heidi.Boarded = true

	//create new bus
	bus := Bus{heidi} //heidi is in the FrontSeat
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the FirstSeat")

	fmt.Println("+++++++++ EXERCISE ++++++++++++")

	rect1 := Rectangle{5, 2}
	area, perimeter := calculateRectangleAreaPerimeter(rect1)
	fmt.Println("The area and Perimeter of Rectangle with length", rect1.Length, "and width", rect1.Width, "is :", area, "and", perimeter)
}
