package main

import "fmt"

func main() {
	var (
		j = 5.99
		g = 7
		bt = 23
	)
	v1 := (9 * j + 8 * float64(g))
	v2 := bt / g
	var v3 = float64(bt) / j

	s := 2 * j + 2 * float64(g)
	
	fmt.Println(
		"costApple", j, 
		"costPear", g,
		"Budget", bt)

	fmt.Println(
		"1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", v1,
		"\n2. Скільки груш ми можемо купити?", v2,
		"\n3. Скільки яблук ми можемо купити?", v3)

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")

	if float64(bt) >= s {
		fmt.Println("можемо")
	} else { 
		fmt.Println("не можемо") 
	}
	
}
