package main
import "fmt"

func main()  {
	var totBelanja, potongan float64
	var asesmen bool

	fmt.Scan(&totBelanja, &asesmen)
	if asesmen {
		potongan = totBelanja * 0.35
	} else {
		potongan = 0
	}
	fmt.Println(totBelanja - potongan)  
}