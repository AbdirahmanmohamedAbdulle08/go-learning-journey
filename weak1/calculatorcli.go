package main

import (
	"fmt"
)

func main() {
	var number1 float64
	var number2 float64
	var operation string
	var result float64
	var again string
	var discard string
	fmt.Println("==============================")
	fmt.Println("     GO CLI CALCULATOR        ")
	fmt.Println("==============================")

	for {
		fmt.Print("\nEnter first number: ")
		_, err := fmt.Scan(&number1)
		if err != nil {
			fmt.Println("Error: Enter correct number!")
			fmt.Scanln(&discard)
			continue
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		_, err = fmt.Scan(&operation)
		if err != nil || (operation != "+" && operation != "-" && operation != "*" && operation != "/") {
			fmt.Println("Invalid operator. Please use +, -, *, or /.")
			fmt.Scanln(&discard)
			continue
		}

		fmt.Print("Enter second number: ")
		_, err = fmt.Scan(&number2)
		if err != nil {
			fmt.Println("Error: enter correct number ")
			fmt.Scanln(&discard)

			continue
		}

		switch operation {
		case "+":
			result = number1 + number2
			fmt.Println("Your result is:", number1, "+", number2, "=", result)
		case "-":
			result = number1 - number2
			fmt.Println("Your result is:", number1, "-", number2, "=", result)
		case "*":
			result = number1 * number2
			fmt.Println("Your result is:", number1, "*", number2, "=", result)
		case "/":
			if number2 == 0 {
				fmt.Println("Error: Lama qaybin karo eber (0)!")
			} else {
				result = number1 / number2
				fmt.Println("Your result is:", number1, "/", number2, "=", result)
			}
		}

		fmt.Print("\nDo you want to calculate again? (y/n): ")
		fmt.Scan(&again)

		if again != "y" && again != "Y" {
			fmt.Println("Thank you for using Go Calculator!")
			break
		}
	}
}
