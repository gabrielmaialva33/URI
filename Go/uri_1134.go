package main

import (
	"fmt"
)

func main() {
	var x, a, g, d = 0, 0, 0, 0
	for {
		fmt.Scanf("%d", &x)
		if x > 0 && x <= 3 {
			if x == 1 {
				a += 1
			} else if x == 2 {
				g += 1
			} else {
				d += 1
			}
		} else if x == 4 {
			break
		} else {
			continue
		}
	}
	fmt.Printf("MUITO OBRIGADO\n")
	fmt.Printf("Alcool: %d\nGasolina: %d\nDiesel: %d\n", a, g, d)

}
