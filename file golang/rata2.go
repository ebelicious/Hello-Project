package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Print(float64(a+b+c+d+e) / float64(5))
}
