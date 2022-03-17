package main

import "fmt"

//M.Ihsan Ramadhan 1303213046
func main() {

	var n, a, r, t int

	fmt.Scanln(&n, &a, &r)

	fmt.Print("0")

	var i int = 1
	for i <= n {

		t = a * i * r
		fmt.Print(" + ", t)
		i++
	}
}
