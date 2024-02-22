package main
import "fmt"

func main()  {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)
	if a >= 0 {
		fmt.Println("terbuka ke atas")
	} else {
		fmt.Println("terbuka ke bawah")
	}
}