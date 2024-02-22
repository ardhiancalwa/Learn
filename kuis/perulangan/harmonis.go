package main
import "fmt"

func main(){
	var  n, angka, R, penyebut, i float64
	
	fmt.Scan(&n)
	angka = 0
	R = 0
	penyebut = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&angka)
		penyebut += 1 / angka
	}
	R = n / penyebut
	fmt.Printf("%.2f", R)
}