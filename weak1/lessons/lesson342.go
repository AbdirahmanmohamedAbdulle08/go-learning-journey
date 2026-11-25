//part two function with variable to type local variable & global variable
//local variable is variable that is declared inside a function and can only be accessed within that function. A global variable is a variable that is declared outside of any function and can be accessed from any function in the program.
//global variable is declared outside of any function and can be accessed from any function in the program. A local variable is declared inside a function and can only be accessed within that function.

//locall variable
package main 
import "fmt"
var country string = "Somalia"
func myname(){
	name:= "abdirahman mohamed"
	fmt.Println("my name is", name)
}
func showCountry(){
	fmt.Println("I live in", country)
}	
func main() {
	myname()
	showCountry()
}