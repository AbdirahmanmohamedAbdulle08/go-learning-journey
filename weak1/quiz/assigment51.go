package main

import (
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("Enter your revenue: ")
	_, err := fmt.Scanln(&revenue)
	if err != nil || revenue <= 0 {
		fmt.Println("Error: Invalid input. Revenue must be a positive number greater than 0.")
		os.Exit(1)
	}
	fmt.Print("Enter your expenses: ")
	_, err = fmt.Scanln(&expenses)
	if err != nil || expenses < 0 {
		fmt.Println("Error: Invalid input. Expenses cannot be a negative number.")
		os.Exit(1)
	}
	fmt.Print("Enter your tax rate (%): ")
	_, err = fmt.Scanln(&taxRate)
	if err != nil || taxRate < 0 || taxRate > 100 {
		fmt.Println("Error: Invalid input. Tax rate must be between 0 and 100.")
		os.Exit(1)
	}
	EBT := revenue - expenses
	Tax := EBT * taxRate / 100
	EAT := EBT - Tax
	// Prevent division by zero if EAT is 0
	if EAT == 0 {
		fmt.Println("Error: EAT (Earnings After Tax) is 0. Cannot compute ratio (division by zero).")
		os.Exit(1)
	}
	ratio := EBT / EAT
	fmt.Printf("\n--- Calculation Results ---\n")
	fmt.Printf("EBT: %.2f\n", EBT)
	fmt.Printf("EAT: %.2f\n", EAT)
	fmt.Printf("Ratio: %.2f\n", ratio)
	resultData := fmt.Sprintf("Revenue: %.2f\nExpenses: %.2f\nTax Rate: %.2f%%\nEBT: %.2f\nEAT: %.2f\nRatio: %.2f\n",
		revenue, expenses, taxRate, EBT, EAT, ratio)
	err = os.WriteFile("financial_results.txt", []byte(resultData), 0644)
	if err != nil {
		fmt.Println("Error: Failed to save results to file:", err)
		os.Exit(1)
	}
	fmt.Println("\nSuccess! Results have been saved to 'financial_results.txt'.")
}
