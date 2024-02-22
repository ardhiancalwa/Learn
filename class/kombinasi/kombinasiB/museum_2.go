package main

import "fmt"

func main() {
	var pengunjung, jumPengunjung, hari int
	var i int

	hari = 1
	for i = 1; i <= 5; i++ {
		fmt.Print("Hari ", hari, " : ")
		fmt.Scan(&pengunjung)
		for pengunjung < 0 || pengunjung > 200 {
			fmt.Print("Hari ", hari, " : ")
			fmt.Scan(&pengunjung)
		}
		jumPengunjung += pengunjung
		hari++
	}
	fmt.Println("Jumlah pengunjung :", jumPengunjung)
}
