package main

import "fmt"
func calculateProfit()float64{
revenue := 500.0
expenses := 150.0
profit:=revenue-expenses
return profit
}
func main(){ 
out:=calculateProfit()
fmt.Print(out)

}
