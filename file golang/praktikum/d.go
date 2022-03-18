package main

import "fmt"

//1303213046_M Ihsan Ramadhan
func main() {
	var s string

	fmt.Scan(&s)
	for s != "selesai" {
		fmt.Println(s)
		fmt.Scan(&s)
	}
}
