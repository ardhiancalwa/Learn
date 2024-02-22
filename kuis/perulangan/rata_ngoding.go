package main
import "fmt"

func main()  {
	var n, totalJam, jam int
	var rata_rata float64
	
	fmt.Scan(&jam)
	n = 0
	totalJam = 0
	for jam >= 0{
		totalJam += jam
		n++
		fmt.Scan(&jam)
	}
	rata_rata = float64(totalJam) / float64(n)
	fmt.Printf("%.2f", rata_rata)
}
