//limitation 
package main
import ("bufio"
"fmt"
"os"
"strings")	

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your FullName")
	fullname, _ := reader.ReadString('\n')
	fullname = strings.TrimSpace(fullname)
	fmt.Println("Your FullName is:", fullname)
}