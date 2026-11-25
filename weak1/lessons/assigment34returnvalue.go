package main

import "fmt"
func getname() string {
	name := "abdirahman"
	return name
}
func getage() int {
	agestd := 21
	return agestd
}
func getscore() float64{
	scorestd := 85.50
	return scorestd
}
var country string
var university string 

func showcounrty(){
	country="somalia"
	fmt.Println(country)
}

func showuniversity(){
	university="Hormuud university"
	fmt.Println(university)
}
func main() {
        a:=getname()
	dada := getage()
        b:=getscore()
	fmt.Println(a)
	fmt.Println(dada)
	fmt.Println(b)
        showcounrty()
        showuniversity()

}
