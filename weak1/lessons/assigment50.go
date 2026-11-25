package main

import "fmt"
func getaddres(){

}
func checkAge(age int) string {
	if age < 18 {
		panic("acces denied")
	} else {
		return "Access granted"
	}
}
func main() {
	var result string
	result = checkAge(20)
	fmt.Print(result)
}
