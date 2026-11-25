//infinate loop is a loop that never ends, it will keep running until the program is terminated.
//break statement is used to terminate the loop or switch statement and transfer execution to the statement immediately following the loop or switch.
//continue statement is used to skip the current iteration of the loop and move to the next iteration.
package main
import "fmt"
func main(){
	// for{
	// 	fmt.Println("this is an infinate loop")
	
	// 
	fmt.Print("this is break")
	for i:=1;i<=10;i++{
		if i==5{
			break
		}
		fmt.Println(i)
}
fmt.Print("this is continue")
for j:=0;j<=5;j++{
	if j==3{
		continue
	}
	fmt.Print(j)

}
}