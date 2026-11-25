//create format string for the lesson31.go file
package main
import "fmt"
func main() {
	name :="abdirahman mohamed"
	age :=20
	message := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(message)
}