package main

import "fmt"

//1303213046_M Ihsan Ramadhan
func main() {
	var nilai, jumlah, n int
	var rata2 float64

	fmt.Scan(&nilai)
	n = 0
	jumlah = 0
	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scan(&nilai)
	}
	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah) / float64(n)
		fmt.Println(rata2)
	}
}
