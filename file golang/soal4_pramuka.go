package main

import "fmt"

//M. Ihsan Ramadhan

func main() {
	var n int
	var b1, b2, b3, b4 bool

	fmt.Scanln(&n)

	var i int = 0
	status := true
	for i < n && status {
		fmt.Scanln(&b1, &b2, &b3, &b4)
		status = b1 && b2 && b3 && b4
		i++
	}

	fmt.Println(status)
}
