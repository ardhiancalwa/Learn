package main
import "fmt"

func main()  {
	var makanan, minuman int
	var tip bool

	fmt.Scan(&makanan, &minuman, &tip)
	if tip {
		fmt.Println(makanan + minuman + 5000)
	} else {
		fmt.Println(makanan + minuman)
	}
}