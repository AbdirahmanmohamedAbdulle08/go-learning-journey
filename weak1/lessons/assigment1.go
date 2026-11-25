// assignment solution to calculate EBT, EAT and ratio of EBT to EAT
package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("Enter your revenue: ")
	fmt.Scanln(&revenue)
	fmt.Print("Enter your expenses: ")
	fmt.Scanln(&expenses)
	fmt.Print("Enter your tax rate (%): ")
	fmt.Scanln(&taxRate)
	EBT := revenue - expenses
	Tax := EBT * taxRate / 100
	EAT := EBT - Tax
	ratio := EBT / EAT
	fmt.Printf("EBT: %.2f\n", EBT)
	fmt.Printf("EAT: %.2f\n", EAT)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
