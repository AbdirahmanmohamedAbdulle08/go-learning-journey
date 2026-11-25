//constant what is a constant in go constants are immutable values which are known at compile time and do not change for the life of the program. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the := syntax.
package main
import "fmt"
func main(){
	const appname="learn Golingo"
	const version=1

	age := 21
	age =90
	// appname="manoow"
	fmt.Println("App name:",appname,"\nversion:",version ,"\nage:",age)
}