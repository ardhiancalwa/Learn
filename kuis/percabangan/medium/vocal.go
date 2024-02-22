package main
import "fmt"

func main() {
	var kata, hurufAwal, hurufAkhir string
	var vocalAwal, vocalAkhir bool

	fmt.Scan(&kata)
	hurufAwal = string(kata[0])
	hurufAkhir = string(kata[len(kata)-1])

	vocalAwal = hurufAwal == "a" || hurufAwal == "i" || hurufAwal == "u" || hurufAwal == "e" || hurufAwal == "o"
	vocalAkhir = hurufAkhir == "a" || hurufAkhir == "i" || hurufAkhir == "u" || hurufAkhir == "e" || hurufAkhir == "o"

	if vocalAwal && vocalAkhir {
		fmt.Println("vokal di awal dan di akhir")
	} else if vocalAwal {
		fmt.Println("vokal di awal")
	} else if vocalAkhir {
		fmt.Println("vokal di akhir")
	} else {
		fmt.Println("tidak ada vokal")
	}
}
