package main
import "fmt"

func main()  {
	var hari, denda int

	fmt.Scan(&hari)
	if hari >= 1 && hari <= 10 {
		denda = hari * 2500
		fmt.Println(denda)
	} else {
		denda = 25000 + (hari * 10) * 5000
		fmt.Println(denda)
	}
}