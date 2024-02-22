package main
import "fmt"

func main()  {
	var hargaMakan int

	fmt.Scan(&hargaMakan)
	if hargaMakan < 5000 {
		fmt.Println("tidak ada makanan dengan harga", hargaMakan)
	} else if hargaMakan >= 5000 && hargaMakan <= 10000 {
		fmt.Println("murah")
	} else if hargaMakan > 10000 && hargaMakan <= 20000 {
		fmt.Println("sedang")
	} else if hargaMakan > 20000 {
		fmt.Println("mahal")
	}
}