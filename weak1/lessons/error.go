package main
import "fmt"
func main(){
var age int
var err error
fmt.Print("Enter your age ")
_, err = fmt.Scan(&age)

if err != nil {
    fmt.Println("invalid age")
} else {
    fmt.Println("your age is", age)
}

}