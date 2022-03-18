package main

import "fmt"

func main() {
	var n, dewasa, kecil, sisa int
	fmt.Scan(&n)
	dewasa = 0
	kecil = 0
	for n > 0 {
		if dewasa != 3 {
			n = n - 5
			dewasa++
		}
		if dewasa == 3 {
			n = n - 2
			kecil++
		}
		if kecil == 5 {
			sisa = n
		}
		if n == 0 {
			break
		}
	}
	if (dewasa > 0) && (kecil == 0) && (sisa == 0) {
		fmt.Print("Dewasa ", dewasa)
	}
	if (dewasa > 0) && (kecil > 0) && (sisa == 0) {
		fmt.Print("Dewasa ", dewasa, " Kecil ", kecil)
	}
	if (dewasa > 0) && (kecil > 0) && (sisa > 0) {
		fmt.Print("Dewasa ", dewasa, " Kecil ", kecil, ", dan ", sisa, " tidak berangkat")
	}
}
