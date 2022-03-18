package main

import "fmt"

//1303213046_M Ihsan Ramadhan

func hitungvolume(r, t float64) float64 {
	pi := 3.14
	return r * r * pi * t
}
func main() {
	var r, t float64
	fmt.Scanln(&r, &t)
	fmt.Println(hitungvolume(r, t))

}
