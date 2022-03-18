package main

import "fmt"

//M. Ihsan Ramadhan 1303213046
func dectobin(num int) { //fungsinya untukkonversi bilangan desimal ke biner
	var biner []int //bikin array untuk menyimpan biner nya

	for num != 0 { //looping untuk membagi 2 dan mengambil sisa bagi, angka yg bisa disimpan cuma 0 dan 1
		biner = append(biner, num%2)
		num = num / 2 //untuk nambahin hasil mod 2 dan hasil bagi 2 ke array biner
	}
	if len(biner) == 0 { //kalo array biner nya kosong, dia akan nge print 0
		fmt.Print("0") //perintah untuk ngeprint 0
	} else {
		for i := len(biner) - 1; i >= 0; i-- { //kalo di array ada nilai selain 0, dia bakal run loop array dari urutan belakang
			fmt.Printf("%d", biner[i]) //i adalah hasil yg berkurang dari loop dan akan diassign ke index array biner
		}
	}
}
func main() { //untuk menginput bilangan desimal dan output biner nya
	var desimal int    //deklarasi variabel desimal adalah integer
	fmt.Scan(&desimal) //input desimalnya
	dectobin(desimal)
}
