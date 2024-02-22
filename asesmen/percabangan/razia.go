package main
import "fmt"

func main()  {
	var nim, plat, platAkhir int
	
	fmt.Scan(&nim, &plat)
	platAkhir = plat%10

	if nim%2 == 0 {
		fmt.Println("NIM genap.")
	} else {
		fmt.Println("NIM ganjil.")
	}
	if platAkhir%2 == 0 {
		fmt.Println("Plat nomor genap.")
	} else {
		fmt.Println("Plat nomor ganjil")
	}
	if (nim%2 == 0 && platAkhir%2 == 0) || (nim%2 != 0 && platAkhir%2 != 0) {
		fmt.Println("Silahkan lewat.")
	} else {
		fmt.Println("Tidak diperbolehkan lewat.")
	}
}
