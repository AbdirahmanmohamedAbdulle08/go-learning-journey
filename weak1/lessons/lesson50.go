package main

import "fmt"

func main() {
	fmt.Println("Program started")

	panic("Something went wrong")

	fmt.Println("Program finished")
}
