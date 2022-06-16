package main

import "fmt"

func main() {
	var (
		applePrice float64 = 5.99
		pearPrice float64 = 7
		Budget float64 = 23
	)
	v1 := 9 * applePrice + 8 * pearPrice
	v2 := Budget / pearPrice
	v3 := Budget / applePrice

	s := 2 * applePrice + 2 * pearPrice
	
	fmt.Println(
		"costApple:", applePrice,
		"costPear:", pearPrice,
		"Budget:", Budget)

	fmt.Println(
		"1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", v1,
		"\n2. Скільки груш ми можемо купити?", v2,
		"\n3. Скільки яблук ми можемо купити?", v3)

	fmt.Print("4. Чи ми можемо купити 2 груші та 2 яблука?")

	if Budget >= s {
		fmt.Print(" - можемо")
	} else { 
		fmt.Print(" - не можемо") 
	}
	
}
