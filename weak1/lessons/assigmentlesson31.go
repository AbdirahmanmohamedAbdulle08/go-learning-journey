//lesson assigment 31 using fmt.Sprintf to create a formatted string
package main
import "fmt"
func main(){
name :="abdirahman mohamed"
age  :=21
country :="somalia"
message := fmt.Sprintf("my name is %s and i am %d year old and my country is %s",name,age,country)
fmt.Print(message)
}