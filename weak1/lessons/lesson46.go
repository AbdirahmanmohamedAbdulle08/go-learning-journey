package main

import "fmt"

func main() {
	var choice int

	fmt.Println("1. Check Balance")
	fmt.Println("2. Send Money")
	fmt.Println("3. Buy Airtime")

	fmt.Print("Dooro adeegga: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Waxaad dooratay Check Balance")
	case 2:
		fmt.Println("Waxaad dooratay Send Money")
	case 3:
		fmt.Println("Waxaad dooratay Buy Airtime")
	default:
		fmt.Println("Doorasho khaldan")
	}
}
