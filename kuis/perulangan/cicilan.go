package main
import "fmt"

func main()  {
	var i, cicilan, bulan, totalCicilan int

	fmt.Scan(&cicilan, &bulan)
	i = 1
	for i <= bulan {
		totalCicilan = cicilan * bulan
		i++
	}
	fmt.Println(totalCicilan)
}