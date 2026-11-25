// EVCPlus Somalia - Condition statements (if, else if, else)
package main

import "fmt"

func main() {
	var haraaga float64 = 500.00
	var ussd string
	for i := 0; i < 3; i++ {
	

	fmt.Println("Garaac koodhka adeegga (*770#):")
	fmt.Scanln(&ussd)

	if ussd == "*770#" {
		fmt.Println("-EVCPLUS-")
		fmt.Println("Fadlan gali PIN-kaaga")
		var pin int
		fmt.Scanln(&pin)

		if pin == 1234 {
			fmt.Println("1. Itus Haraaga")
			fmt.Println("2. Lacag Wareeji")
			fmt.Println("3. Ku hadal ku shubo")

			var option int
			fmt.Scanln(&option)

			if option == 1 {
				fmt.Printf("<-EVCPlus-> Haraagaagu waa $%.2f\n", haraaga)
			} else if option == 2 {
				var number string
				var amount float64
				var confirm int

				fmt.Println("Fadlan gali Mobile-ka:")
				fmt.Scanln(&number)
				fmt.Println("Fadlan gali lacagta:")
				fmt.Scanln(&amount)

				fmt.Printf("Mahubtaa inaad $%.2f u wareejiso %s?\n1. Haa\n2. Maya\n", amount, number)
				fmt.Scanln(&confirm)

				if confirm == 1 {
					if amount > haraaga {
						fmt.Println("Haraaga xisaabtaadu kuguma filna.")
					} else if amount <= 0 {
						fmt.Println("Fadlan gali lacag sax ah.")
					} else {
						haraaga -= amount
						fmt.Printf("[-EVCPlus-] Waxaad $%.2f u wareejisay %s. Haraagaagu waa $%.2f\n", amount, number, haraaga)
					}
				} else {
					fmt.Println("Waad ka noqotay lacag wareejinta.")
				}
			} else if option == 3 {
				var airtime float64
				var confirm int

				fmt.Println("Fadlan gali lacagta ku hadalka:")
				fmt.Scanln(&airtime)

				fmt.Printf("Mahubtaa inaad ku shubato $%.2f ku hadal ah?\n1. Haa\n2. Maya\n", airtime)
				fmt.Scanln(&confirm)

				if confirm == 1 {
					if airtime > haraaga {
						fmt.Println("Haraaga xisaabtaadu kuguma filna.")
					} else if airtime <= 0 {
						fmt.Println("Fadlan gali lacag sax ah.")
					} else {
						haraaga -= airtime
						fmt.Printf("[-EVCPlus-] Waxaad ku shubatay $%.2f ku hadal ah. Haraagaagu waa $%.2f\n", airtime, haraaga)
					}
				} else {
					fmt.Println("Waad ka noqotay ku hadal ku shubashada.")
				}
			} else {
				fmt.Println("Fadlan dooro lambar sax ah (1 ilaa 3).")
			}
		} else {
			fmt.Println("PIN-kaagu waa khalad.")
		}
	} else {
		fmt.Println("Koodhka aad garaacday waa khalad. Fadlan garaac *770#")
	}
}
}
