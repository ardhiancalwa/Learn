// Buatlah fungsi saja
package main

import "fmt"

func totalTabungan(x, y, z int) int {
	var i, total_tabungan, tabungan_akhir int
	total_tabungan = 0
	tabungan_akhir = 0
	for i = 1; i <= z; i++ {
		if i%2 == 0 || i%3 == 0 {
			total_tabungan = total_tabungan + x
			tabungan_akhir = total_tabungan
            
		}
	}
	return tabungan_akhir + y
}

func main() {
	var hasil, a, b, c int
	fmt.Scan(&a, &b, &c)
	hasil = totalTabungan(a, b, c)
	fmt.Println(hasil)
}