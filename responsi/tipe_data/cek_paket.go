package main
import "fmt"

func main()  {
	var panjang, lebar, tinggi, berat, volume int
	var check bool

	fmt.Scan(&panjang, &lebar, &tinggi, &berat)
	volume = panjang * lebar * tinggi
	check = volume <= 1000000 && berat <= 30000
	fmt.Println(check)
}