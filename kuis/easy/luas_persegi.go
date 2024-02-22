package main
import "fmt"

func main() {
	var panjang, lebar, luas float64

	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	fmt.Println(luas)
}