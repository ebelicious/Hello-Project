package main

import "fmt"

//1303213046_M Ihsan Ramadhan
func main() {
	var n, jumlah, t1, t2, t3 int
	fmt.Scanln(&n)
	jumlah = 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlah++
		}
	}
	fmt.Println(jumlah)
}
