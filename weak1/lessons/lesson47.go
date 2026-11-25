package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	var qoraal string

	fmt.Println("1. Check Balance")
	fmt.Println("2. Send Money")
	fmt.Println("3. Buy Airtime")

	fmt.Print("Dooro adeegga: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		qoraal = "Waxaad dooratay Check Balance"
		fmt.Println(qoraal)

	case 2:
		qoraal = "Waxaad dooratay Send Money"
		fmt.Println(qoraal)

	case 3:
		qoraal = "Waxaad dooratay Buy Airtime"
		fmt.Println(qoraal)

	default:
		qoraal = "Doorasho khaldan"
		fmt.Println(qoraal)
	}

	err := os.WriteFile("switchsae.txt", []byte(qoraal), 0644)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Successfully saved!")
}
