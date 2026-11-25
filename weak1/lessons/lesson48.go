package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("students.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
