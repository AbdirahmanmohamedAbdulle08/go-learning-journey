package main

import "fmt"

func main() {
    price := 19.5678
    tax := 2.4567

    total := price + tax

    fmt.Printf("Price: %.2f\n", price)
    fmt.Printf("Tax: %.2f\n", tax)
    fmt.Printf("Total: %.2f\n", total)
}