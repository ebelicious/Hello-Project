package main

import "fmt"

func main() {
	var n, apb, apk, sel int

	fmt.Scanln(&n)
	i := n
	apb = 3
	apk = 5

	for {
		i--
		sel = n - 5
		for sel > 0 && apb != 0 {
			if sel > 0 {
				apb = apb - 1
			}
			for sel > 0 && apb == 0 {
				if sel > 0 {
					apk = apk - 1
				}
			}
		}
		if i == 0 {
			break
		}
	}
	fmt.Println("besar = ", apb, "kecil = ", apk)
}
