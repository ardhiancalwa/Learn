package main
import "fmt"

func main() {
	var unit, biayaTotal float64

	fmt.Scan(&unit)
	if unit < 200 {
		biayaTotal = unit * 12
	} else if unit >= 200 && unit < 400 {
		biayaTotal = unit * 15
	} else if unit >= 400 && unit < 600 {
		biayaTotal = unit * 18
	} else if unit >= 600 {
		biayaTotal = unit * 20
	}

	if biayaTotal > 400 {
		biayaTotal += biayaTotal * 0.2
	} else if biayaTotal < 100 {
		biayaTotal = 100
	}
	fmt.Println(biayaTotal)
}