package main

import "fmt"

func dapatKTP(umur int, golang bool) bool {
	// mengembalikan boolean true jika umur lebih besar dari atau sama dengan 10
	// dan fasih berbahasa golang
	var isKTP bool
	if umur >= 10 && golang {
		isKTP = true
	} else {
		isKTP = false
	}
	return isKTP

}

func main() {
	var umur int
	var golang bool
	fmt.Scan(&umur, &golang)
	fmt.Println(dapatKTP(umur, golang))
}
