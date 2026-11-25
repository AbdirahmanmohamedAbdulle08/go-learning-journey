package main
import "fmt"
func main(){
name :="abdirahman mohamed"
age  :=21
country :="somalia"
message := fmt.Sprintf("my name is %s\n i am %d year old\n my country is %s",name,age,country)
fmt.Print(message)
}