package main

import "fmt"

func main() {
	var baca, cpek, gopek, srebu, dwarebu int
	fmt.Scanln(&baca)
	if baca == 100 {
		cpek++
		fmt.Println("Rp.100 = ", cpek, "keping")
	}
	if baca == 500 {
		gopek++
		fmt.Println("Rp.500 = ", gopek, "keping")
	}
	if baca == 1000 {
		srebu++
		fmt.Println("Rp.1000 = ", srebu, "keping")
	}
	if baca == 2000 {
		dwarebu++
		fmt.Println("Rp.2000 = ", dwarebu, "keping")
	}
	i := 1
	for (baca == 100) || (baca == 500) || (baca == 1000) || (baca == 2000) {
		i++
		fmt.Scanln(&baca)
		if baca == 100 {
			cpek++
			fmt.Println("Rp.100 = ", cpek, "keping")
		}
		if baca == 500 {
			gopek++
			fmt.Println("Rp.500 = ", gopek, "keping")
		}
		if baca == 1000 {
			srebu++
			fmt.Println("Rp.1000 = ", srebu, "keping")
		}
		if baca == 2000 {
			dwarebu++
			fmt.Println("Rp.2000 = ", dwarebu, "keping")
		}
	}
}
