//user input improve
package main
import ("fmt" 
"bufio" 
"os"
 "strings")

func main() {
	var name string
	var age int
	var country string
reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to student registration system")
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter your age: ")
	fmt.Fscanf(reader, "%d\n", &age)
	fmt.Print("Enter your country: ")
	country, _ = reader.ReadString('\n')
	country = strings.TrimSpace(country)
	fmt.Println("Student Information:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
}	