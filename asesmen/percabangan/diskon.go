package main
import "fmt"

func main() {
	var beli int
	var totalAkhir float64

	fmt.Scan(&beli)
	totalAkhir = float64(beli)
	if beli >= 275000 {
		totalAkhir = float64(beli) - (float64(beli) * 5 / 100)
	}
	fmt.Println(totalAkhir)
}
