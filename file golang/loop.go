package main

import "fmt"

//M Ihsan Ramadhan 1303213046

func main() {
	var x, y int
	fmt.Scanln(&y)
	x = 1

	for x*2 <= y {
		x = x * 2
	}
	for x >= 1 {
		fmt.Print(y / x)
		y = y % x
		x = x / 2
	}
}
