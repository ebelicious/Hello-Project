package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	nama := name.Text()
	fmt.Println("Nama Saya Adalah : ", nama)
}
