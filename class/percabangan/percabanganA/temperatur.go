package main
import "fmt"

func main()  {
	var temp1, temp2, temp3, temp4, temp5 float64
	var sensor string

	fmt.Scan(&temp1, &temp2, &temp3, &temp4, &temp5)
	if temp1 < temp2 && temp2 < temp3 && temp3 < temp4 && temp4 < temp5 {
		sensor = "Stabil naik"
	} else if temp1 > temp2 && temp2 > temp3 && temp3 > temp4 && temp4 > temp5 {
		sensor = "Stabil turun"
	} else {
		sensor = "Tidak stabil"
	}
	fmt.Println(sensor)
}