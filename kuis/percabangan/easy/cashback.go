package main
import "fmt"

func main()  {
	var totBelanja, cashback int
	var kartu bool

	fmt.Scan(&totBelanja, &kartu)
	if totBelanja >= 500000 && kartu {
		cashback = totBelanja * 20/100
		fmt.Println(cashback)
	} else {
		cashback = 0
		fmt.Println(cashback)
	}
}