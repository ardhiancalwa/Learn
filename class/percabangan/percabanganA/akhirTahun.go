package main
import "fmt"

func main()  {
	var  kartu, diskon, cashback, sedia bool
	var totBelanja int

	fmt.Scan(&totBelanja, &sedia)
	kartu = sedia == true
	diskon = totBelanja >= 100000
	cashback = totBelanja >= 200000 && kartu

	if cashback == true {
		totBelanja = (totBelanja - (totBelanja * 10 / 100)) - 75000
	} else if diskon == true {
		totBelanja = totBelanja - (totBelanja * 10 / 100)
	}
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp.", totBelanja)
}