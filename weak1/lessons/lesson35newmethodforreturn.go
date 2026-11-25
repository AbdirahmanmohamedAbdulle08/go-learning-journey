package main 

import "fmt"

func add()(result int){
result=5+4
return
}

func main(){
out:=add()
fmt.Print(out)
}