package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	var age int
	var course string

	read := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome Student Form")

	fmt.Print("Enter your full name: ")
	fullname, err := read.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	name = strings.TrimSpace(fullname)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your course: ")
	fmt.Scan(&course)

	qoraal := fmt.Sprintf(
		"Student Name: %s\nAge: %d\nCourse: %s\n",
		name, age, course,
	)

	err = os.WriteFile("students.txt", []byte(qoraal), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Successfully saved!")
}
