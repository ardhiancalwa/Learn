package main
import "fmt"

func main()  {
	var kendaraan string
	var lama, tarifAwal, tarifBerikutnya, totTarif int
	fmt.Scan(&kendaraan, &lama)
	if kendaraan == "motor" {
		tarifAwal = 2000
		tarifBerikutnya = 500
	} else if kendaraan == "mobil" {
		tarifAwal = 3000
		tarifBerikutnya = 1000
	}
	if lama > 1 {
		totTarif = tarifAwal + tarifBerikutnya * (lama-1)
	}
	fmt.Println(totTarif)
}