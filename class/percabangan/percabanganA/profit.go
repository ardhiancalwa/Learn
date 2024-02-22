package main
import "fmt"

func main()  {
	var profitNow, profitBefore, jumProfit float64
	var keuntungan string

	fmt.Scan(&profitNow, &profitBefore)
	if profitNow == profitBefore {
		keuntungan = "Tetap"
		fmt.Println(keuntungan)
	} else if profitNow < profitBefore {
		jumProfit = profitBefore - profitNow
		keuntungan = "Peningkatan sebesar"
		fmt.Println(keuntungan, jumProfit)
	} else if profitNow > profitBefore {
		jumProfit = profitNow - profitBefore
		keuntungan = "Penurunan sebesar" 
		fmt.Println(keuntungan, jumProfit)
	}
}