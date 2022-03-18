package main

import "fmt"

func main() {
	var a, b, c, d int
	var rata float64
	fmt.Scan(&a, &b, &c, &d)

	rata = float64(a+b+c+d) / 4

	if rata >= 3.50 {
		fmt.Println("bagus")
	}
	if rata <= 1.50 {
		fmt.Println("kurang")
	}
	if (rata > 1.50) && (rata < 3.50) {
		fmt.Println("sedang")
	}
}
