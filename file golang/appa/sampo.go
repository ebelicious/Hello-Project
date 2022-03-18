package main

import "fmt"

func main() {
	var sampo, ujan string
	fmt.Scan(&sampo, &ujan)
	if sampo == "tidak" {
		fmt.Println("tidak pergi ke supermarket")
	}
	if (sampo == "ya") && (ujan == "ya") {
		fmt.Println("tidak pergi ke minimarket")
	}
	if (sampo == "ya") && (ujan == "tidak") {
		fmt.Println("pergi ke minimarket")
	}
}
