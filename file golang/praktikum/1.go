package main

import "fmt"

//M Ihsan R 1303213046

func main() {
	var x1, x2, y1, y2 int
	var gradient float64
	fmt.Scanln(&x1, &y1, &x2, &y2)
	gradient = float64(y1-y2) / float64(x1-x2)
	fmt.Println(gradient)
	fmt.Scanln(&x1, &y1, &x2, &y2)
	gradient = float64(y1-y2) / float64(x1-x2)
	fmt.Println(gradient)
	fmt.Scanln(&x1, &y1, &x2, &y2)
	gradient = float64(y1-y2) / float64(x1-x2)
	fmt.Println(gradient)
}
