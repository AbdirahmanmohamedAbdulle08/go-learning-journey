package main
import ("bufio"
"os"
"strings"
"fmt")
func main(){
read := bufio.NewReader(os.Stdin)
fmt.Println("Enter your full name")
fullname,_:=read.ReadString('\n')
fullname =strings.TrimSpace(fullname)
fmt.Println("your fullname is ",fullname)
}