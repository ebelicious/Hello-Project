package main

import "fmt"

//M IHSAN R 1303213046

func main() {
	var A, B, C string
	var hasil bool
	fmt.Scanln(&A, &B, &C)
	hasil = A == B || A == C || B == C
	fmt.Println(hasil)
}
