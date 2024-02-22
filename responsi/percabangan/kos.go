package main
import "fmt"

func main()  {
	var panjang, lebar, harga int
	var tipe string

	fmt.Scan(&panjang, &lebar, &tipe, &harga)
	if panjang*lebar >= 30 && tipe == "Cowo" && harga >= 800000 && harga <= 1000000 {
		fmt.Println("Idaman Galih")
	} else {
		fmt.Println("Bukan idaman Galih")
	}
}