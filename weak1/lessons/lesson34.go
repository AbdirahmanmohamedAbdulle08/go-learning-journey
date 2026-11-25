//funtion with return value
package main
import "fmt"
func addition()int{
	result :=5+6
	return result
}
func main() {
	sum := addition()
	fmt.Println("The sum is:", sum)
}

