package main
import "fmt"

func main()  {
	var hari, bayar int
	var denda string

	fmt.Scan(&hari)
	if hari >= 1 && hari <= 5 {
		bayar = hari * 1000
		fmt.Println(bayar)
	} else if hari > 5 && hari <= 10 {
		bayar = hari * 2000
		fmt.Println(bayar)
	} else {
		denda = "cabut keanggotaan"
		fmt.Println(denda)
	}
}