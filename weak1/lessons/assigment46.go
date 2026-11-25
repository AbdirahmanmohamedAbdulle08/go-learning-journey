// assigment swiitch
package main

import "fmt"

func main() {
	var haraaga float64 = 200.2
	fmt.Println("Welcome evcplus")
	fmt.Println("1.Itus Haraaga")
	fmt.Println("2.Dir Evc-plus")
	fmt.Println("3.ku shubo kuhadal")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Print("Haraaga waa ", haraaga)
	case 2:
		fmt.Println("Fadlan gali numberka lacagta udirayso")
		var number int
		fmt.Scan(&number)
		fmt.Println("Fadlan gali Xadiga lacagta udirayso")
		var lacagta float64
		fmt.Scan(&lacagta)
		fmt.Println("Mahubtaa Inaa Udirayso Numberkaa", number, "lacag dhan", lacagta, "\n1.Haa", "\n2.May")
		var verify int
		fmt.Scan(&verify)
		if verify == 1 {
			for i := 0; i < 3; i++ {
				fmt.Println("fadlan gali pin")
				var pin int
				fmt.Scan(&pin)
				if pin == 1234 {
					if haraaga >= lacagta {
						haraaga -= lacagta
						fmt.Println("Waxaa Lacagtaa", lacagta, "udirtay numberkaan", number, "Haragaagu waa", haraaga)
						break
					} else {
						fmt.Println("Haraagagu Kuguma Filna")
						break
					}
				} else {
					fmt.Print("Pinka Waa Khalad")
				}
			}

		} else {
			return
		}
	default:
		fmt.Print("wax maadan dooran")
	}

}
