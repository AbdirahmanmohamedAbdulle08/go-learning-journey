//lesson for 37 38 39 lear condition statement if else if and else
package main

import "fmt"	
func main(){	
	var haraaga float64=1000.00
	fmt.Println("salaam Bank")
	fmt.Println("1.itus haraaga")
	fmt.Println("2.lacag dhigasho")
	fmt.Println("3.lacag qaadasho")
	fmt.Println("4.ka wareejin evc")
	fmt.Println("5.ka wareejin account")
	fmt.Println("6.hubi wareejinta account")
	fmt.Println("7.mareynta bankiga")
	fmt.Println("8.ka bax")
	var option int
	fmt.Scanln(&option)
	if option==1{
		fmt.Println("lacagta haraaga waa ",haraaga)
	}else if option==2{
		var deposit float64
		fmt.Println("fadlan gali xadigalacag dhigasho")
		fmt.Scanln(&deposit)
		fmt.Println("mahubtaa inaa dhigatay lacagta", deposit, "\n1.haa","\n2.maya")
		var choice int
		fmt.Scanln(&choice)
		if choice==1{
			fmt.Println("lacagta aad dhigatay waa ",deposit, "fadlan gali pinkaaga banka")
			var pin int
			fmt.Scanln(&pin)
			if pin==9999{
				haraaga+=deposit
				fmt.Println("lacagta haraaga waa ",haraaga)
			} else{
				fmt.Println("pinkaaga waa khalad")
			}
		} else if choice==2{
			fmt.Println("waad ka noqotay lacag dhigashada")
		}
		
	} else if option==3{
		var withdraw float64
		fmt.Println("fadlan gali xadiga lacag qaadasho")
		fmt.Scanln(&withdraw)
		fmt.Println("mahubtaa inaa qaadatay lacagta", withdraw, "\n1.haa","\n2.maya")
		var choice int
		fmt.Scanln(&choice)
		if choice==1{
			fmt.Println("lacagta aad qaadatay waa ",withdraw, "fadlan gali pinkaaga banka")
			var pin int
			fmt.Scanln(&pin)
			if pin==9999{
					if withdraw>haraaga{
						fmt.Println("lacagta aad rabto inaad qaadato waa ka badan tahay haraagaaga")
						} else{
							haraaga-=withdraw
							fmt.Println("lacagta haraaga waa ",haraaga)
						}
					}else{
						fmt.Println("pinkaaga waa khalad")}	
			}else if choice==2{
				fmt.Println("waad ka noqotay lacag qaadashada")	
			}
	} else {
		fmt.Println("fadlan dooro option sax ah")
	}
}

	