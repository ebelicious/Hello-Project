package main

import "fmt"

// M IHSAN R 1303213046
func main() {
	var tahun1, tahun2, tahun3, tahun4 int
	var k1, k2, k3, k4 bool

	fmt.Scanln(&tahun1)
	fmt.Scanln(&tahun2)
	fmt.Scanln(&tahun3)
	fmt.Scanln(&tahun4)
	k1 = tahun1%400 == 0 || tahun1%4 == 0 && tahun1%100 != 0
	k2 = tahun2%400 == 0 || tahun2%4 == 0 && tahun2%100 != 0
	k3 = tahun3%400 == 0 || tahun3%4 == 0 && tahun3%100 != 0
	k4 = tahun4%400 == 0 || tahun4%4 == 0 && tahun4%100 != 0
	fmt.Println(k1 && k2 && k3 && k4)
	fmt.Println(k1, k2, k3, k4)
}
