package main

import "fmt"

//1303213046_M Ihsan Ramadhan
func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	if d > a && d > b && d > c {
		fmt.Println("ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}
}
