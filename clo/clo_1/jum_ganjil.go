package main

import "fmt"

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlahGanjil(bil1, bil2))
}

func jumlahGanjil(bil1, bil2 int) int {
	var i, totGanjil int
	if bil1 < bil2 {
		for i = bil1; i <= bil2; i++ {
			if i%2 != 0 {
				totGanjil += i
			}
		}
	} else {
		for i = bil2; i <= bil1; i++ {
            if i%2!= 0 {
                totGanjil += i
            }
        }
	}
	return totGanjil
}
