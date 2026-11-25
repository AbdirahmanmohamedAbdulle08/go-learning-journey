package main
import "fmt"
func getStudentName()(fullname string){
fullname="Abdirahman Mohamed abdulle"
return
}
func calculateTotal() (addition int){
addition = 150 + 250
return 100
}
func main(){
var studentname string =getStudentName()
var calculator =calculateTotal()
fmt.Println(studentname)
fmt.Println(calculator)
}