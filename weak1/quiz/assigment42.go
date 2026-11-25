//if condition use return statement return statement is used to exit the function and return a value to the caller
package main

import "fmt"	
func checkpin(pin int)bool{
	if pin==1234{
		return true
	}else{
		return false
	}
}
func main(){
	var pin int
	fmt.Println("Fadlan gali PIN-kaaga:")
	fmt.Scanln(&pin)
	if checkpin(pin){
		fmt.Println("PIN-kaaga waa sax")
	}else{
		fmt.Println("PIN-kaaga waa khalad")
	}
}