package main 
import "fmt"

func main() {
	var Name string = "Abdirahman Mohamed"
	var age int =21
	var mathscore int=90
	var sciencescore int=85
	var Englishscore int=88
	fmt.Println("Name:", Name)
	fmt.Println("Age:", age)
	fmt.Println("Math Score:", mathscore)
	fmt.Println("Science Score:", sciencescore)
	fmt.Println("English Score:", Englishscore)
	fmt.Println("Total Score:", mathscore+sciencescore+Englishscore)
fmt.Println("Average Score:", (mathscore+sciencescore+Englishscore)/3)	
}	
