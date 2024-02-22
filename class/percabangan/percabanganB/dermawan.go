package main
import "fmt"

func main()  {
	var minggu, uang  int
	var total int

	fmt.Scan(&minggu)
	if minggu == 4 {
		for i := 1; i <= 4; i++ {
			fmt.Scan(&uang)
			total += uang
		}
		fmt.Println(total)
	}
	if minggu == 5 {
		for i := 1; i <= 5; i++ {
			fmt.Scan(&uang)
			total += uang
		}
		fmt.Println(total)
	}

}