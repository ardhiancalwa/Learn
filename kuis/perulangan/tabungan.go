package main
import "fmt"

func main()  {
	var i, tabungan, minggu, tabunganTotal int

	tabunganTotal = 0
	fmt.Scan(&tabungan, &minggu)
	
	for i = 1; i <= minggu; i++ {
		tabunganTotal += tabungan
        tabungan += 2500
	} 
	fmt.Println(tabunganTotal)
}