// multiple line comment
package main
import "fmt"
func main() {
	var message string
	// message ="my name is abdirahman mohamed \ni am 20 year old \ni live in somalia"
	message = `my name is abdirahman mohamed
i am 20 year old
i live in somalia`
	fmt.Println(message)
}