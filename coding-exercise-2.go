package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string  //library member name

type LendAudit struct {
	checkOut time.Time // we utilize the test package from the standard library.
	//It has a structure names Time and that structure tracks whatever time we placed into it. We are going to use current time when we create
	// our program, which will indicate that the book was checked out like right now or checkback in right now
	checkIn time.Time
}

//Next we will have member of a Library
type Member struct {
	name  Name
	books map[Title]LendAudit //books map here we are using a title for a key.
	//So thats the title of the book (Thats the key) so that we can lookup easily any books,
	// The LandAudit is the structure here which has a checkout and check in time
}

//when a member of the library checks out a book, the checkout time will be set to the current time and
//if they havent checked it in. This check in time here , thats going to be zero. If its zero that means the book has not been checked in
// Once we have a valid checkout and checkin time, we know that the book has been checked out and the member and has not yet been checked in

// Once we have a valid checkout and check in time , we know that the book has been checked out and the then returned
// So we can use the logic to determine if a member needs to return a book, because we would know that the checkin time is zero and we can just
// check for that

type BookEntry struct {
	total  int // total owned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

// So the library is going to be able to determine who has what books based on this member variable here and then the total books that it has
// are the book entry structure and so we know total books and how many other lended. This should do for our type aliases and structures

//We are going to create a function which is going to take a pointer to an existing member and then going to iterate through each book
func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String() // we are going to use the .String() on the time function to print it out in string format
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime) // print out the member information that which member is being printed and the current status of their audit
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

// Thats all we need to do to get information on the library and which books have been lended out and how many the library has
//Now we need to make functions to return a book and check out a book

// Have a function for checking out a book first and then we will have a function for returning the book

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title] // we are checking if the title of the book exists in the library
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}
	if book.lended == book.total { // Next we will check if the book has been lended or not
		fmt.Println("No more books available to lend")
		return false
	}

	//At this point we have atleast one book to lend
	book.lended += 1
	library.books[title] = book // reassigning the existing book entries, we are updating it and re assign it back into the library

	//update the member lending
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true

	// The check in time will have a default value of zero

}

// Return the book to the library
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}
	// After this we need to make sure that the person returning the book actually checked it out
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	//Update the library
	book.lended -= 1
	library.books[title] = book // reassigning the existing book entries, we are updating it and re assign it back into the library

	audit.checkIn = time.Now()
	member.books[title] = audit // Reassign the audit to the member information
	return true                 // That the book is returned to the library

}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Lets learn GO"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go BootCamp"] = BookEntry{
		total:  4,
		lended: 0,
	}

	//Add some members
	library.members["John"] = Member{"Name", make(map[Title]LendAudit)}
	library.members["Henry"] = Member{"Henry", make(map[Title]LendAudit)}
	library.members["Casey"] = Member{"Casey", make(map[Title]LendAudit)}

	fmt.Println("\nInitial status of the library")
	printLibraryBooks(&library) // we will take a pointer to the library
	printMemberAudits(&library)

	//Checkout the book
	//Take one of the members or access the number through the library
	member := library.members["John"]
	checkedOut := checkoutBook(&library, "Go BootCamp", &member)
	fmt.Println("\nCheck out a book: ")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go BootCamp", &member)
	fmt.Println("\nCheck in a book: ")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

}
