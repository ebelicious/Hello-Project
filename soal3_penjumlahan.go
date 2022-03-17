package main

import "fmt"

//M IHSAN RAMADHAN 1303213046

func main() {
	var n, bilangan int

	fmt.Scanln(&n) //input

	var i int = 0
	var total = 0
	for i < n {
		fmt.Scanln(&bilangan)
		for bilangan < 0 || bilangan > 9 {
			fmt.Scanln(&bilangan)
		}

		total += bilangan
		i++
	}

	fmt.Println(total) //output
}
