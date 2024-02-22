package main
import "fmt"

func main()  {
	var n, uang, jumlah int

	fmt.Scan(&n)
	jumlah = 0
	if n == 4 {
		for i := 1; i <= n; i++ {
			fmt.Scan(&uang)
			jumlah += uang
		}
	} else if n == 5 {
		for i := 1; i <= n; i++ {
			fmt.Scan(&uang)
			jumlah += uang
		}
	}
	fmt.Println(jumlah)
}