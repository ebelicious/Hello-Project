package main

import "fmt"

//M IHSAN RAMADHAN 1303213046

func main() {

	var N, c1, c2 int

	fmt.Scanln(&N) //asumsikan N = 1234

	c1 = N % 10 //adalah 4
	N = N / 10  //1234 => 123
	konsekutif := true
	for N > 0 && konsekutif {
		c2 = N % 10 // 3
		konsekutif = c1-c2 == 1 || c2-c1 == 1
		c1 = c2
		N = N / 10

	}

	fmt.Println(konsekutif)
}
