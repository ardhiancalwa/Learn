package main
import "fmt"

func main() {
	var i, n, detik, jumlah int
	var rata_rata float64

	fmt.Scan(&n)
	jumlah = 0
	i = 1
	for i <= n {
		fmt.Scan(&detik)
		jumlah += detik
		rata_rata = float64(jumlah) / float64(n)
		i++
	}
	fmt.Println(jumlah)
	fmt.Printf("%.2f", rata_rata)

}
