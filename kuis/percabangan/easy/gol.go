package main
import "fmt"

func main()  {
	var gawangSendiri, gawangLawan int
	var hasilAkhir string

	fmt.Scan(&gawangLawan, &gawangSendiri)
	if gawangLawan > gawangSendiri {
		hasilAkhir = "menang"
	} else if gawangLawan == gawangSendiri {
		hasilAkhir = "draw"
	} else {
		hasilAkhir = "kalah"
	}
	fmt.Println(hasilAkhir)
}