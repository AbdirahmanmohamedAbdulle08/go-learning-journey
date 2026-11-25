package main
import ("bufio"
"os"
"strings"
"fmt")
func main(){
read :=bufio.NewReader(os.Stdin)
fmt.Print("Enter Your name")
fullname,_:=read.ReadString('\n')
fullname =strings.TrimSpace(fullname)
fmt.Print("your name is",fullname)
}