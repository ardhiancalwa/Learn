package main
import "fmt"

func main()  {
	var index float64
	var predikat string

	fmt.Scan(&index)
	if index < 2.00 {
		predikat = "poor"
	} else if index >= 2.00 && index <= 2.75 {
		predikat = "Fair"
	} else if index > 2.75 && index <= 3.00 {
		predikat = "Satisfactory"
	} else if index > 3.00 && index <= 3.50 {
		predikat = "Very Good"
	} else {
		predikat = "Excelent"
	}
	fmt.Println(predikat)
}