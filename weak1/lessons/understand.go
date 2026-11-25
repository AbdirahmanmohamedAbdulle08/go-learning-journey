package main

import "fmt"
func getnumber()int{
 result:=5
 return result
}
func getname()string{
fullname:="myname is abdirahman mohamed abdulle"
return fullname
}
func addition()float64{
number1:=5.0
number2:=6.0
addi :=number1+number2
return addi
}
func main(){ 
number :=getnumber()
name:=getname()
sum:=addition()
fmt.Print(number,"\n")
fmt.Print(name,"\n")
fmt.Println(sum)
}
