package main
import "fmt"

func main()  {
	var warna1, warna2 string
	var sekunder string

	fmt.Scan(&warna1, &warna2)
	if (warna1 == "merah" && warna2 == "kuning") || (warna1 == "kuning" && warna2 == "merah") {
		sekunder = "orange"
	} else if (warna1 == "kuning" && warna2 == "biru") || (warna1 == "biru" && warna2 == "kuning") {
		sekunder = "hijau"
	} else if (warna1 == "biru" && warna2 == "merah") || (warna1 == "merah" && warna2 == "biru") {
		sekunder = "ungu"
	}
	fmt.Println(sekunder)
}