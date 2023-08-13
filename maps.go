package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//Function to print out server status
func printServerStatus(servers map[string]int) {
	//Number of servers , use the len function on the servers map to get the length
	fmt.Println("\nThere are", len(servers), "servers")

	//Number of servers for each status (Online, Offline, Maintenance, Retired)
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status") //panic function will immediately terminate the program with this error message
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are Offline")
	fmt.Println(stats[Maintenance], "servers are under going Maintenance")
	fmt.Println(stats[Retired], "servers are Retired")
}
func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11 // key value
	shoppingList["milk"] = 7
	shoppingList["bread"] += 1 // assigning it to value 0 and then adding 1

	shoppingList["eggs"] += 1 // making one dozen eggs

	fmt.Println(shoppingList)
	//deletes based on the keys

	delete(shoppingList, "milk")
	fmt.Println("Milk will be deleted, new list : ", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"] // if cereal is available in shopping list then it returns found as true else false
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yes", cereal, "boxes")
	}

	//Iterate through the map and count the number of items we have purchased
	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("There are", totalItems, "items on the shopping list")

	fmt.Println("++++++++++++++++++++++++++++++++++++")

	servers := []string{"darkstaar", "aiur", "omicron", "w359", "baseline"}

	serverStatusMap := make(map[string]int)
	// String will be mapped to the servers and int , key will be mapped to the server status
	// we will iterate over the servers slice using range and '_' ignore the index ,and use the values only
	for _, server := range servers {
		serverStatusMap[server] = Online
	}

	//Disply server info
	printServerStatus(serverStatusMap)

	//Change darstar to Retired
	serverStatusMap["darkstaar"] = Retired

	//Change aiur to Offline
	serverStatusMap["aiur"] = Offline

	//Disply server info
	printServerStatus(serverStatusMap)

	//change all server status to 'Maintenance'
	for server, _ := range serverStatusMap {
		serverStatusMap[server] = Maintenance
	}

	//Disply server info
	printServerStatus(serverStatusMap)
}
