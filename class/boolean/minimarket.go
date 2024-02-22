package main
import "fmt"

func main()  {
	var kartu, diskon, cashback, bersedia bool
	var totalBelanja int

	fmt.Scan(&totalBelanja, &bersedia)
	diskon = totalBelanja >= 100000
	kartu = bersedia == true && diskon == true
	cashback = totalBelanja >= 200000 && kartu == true
	fmt.Println("Kartu ?", kartu)
	fmt.Println("Diskon ?", diskon)
	fmt.Println("Cashback ?", cashback)
}