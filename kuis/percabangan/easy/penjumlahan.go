package main
import "fmt"

func main()  {
	var bil1, bil2, bil3, jumlah int
	var hitung string
	fmt.Scan(&bil1, &bil2, &bil3, &jumlah)

	if bil1 + bil2 + bil3 == jumlah {
		hitung = "benar"
	} else {
		hitung = "salah"
	}
	fmt.Println(hitung)
}