package main
import "fmt"

func main()  {
	var pembilang1, penyebut1, pembilang2, penyebut2, jumlah int

	fmt.Scan(&pembilang1, &penyebut1, &pembilang2, &penyebut2)
	if penyebut1 == penyebut2 {
		jumlah = pembilang1 + pembilang2
		fmt.Println(jumlah)
	} else {
		jumlah = 0
		fmt.Println(jumlah)
	}
}