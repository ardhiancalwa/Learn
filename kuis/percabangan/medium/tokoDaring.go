package main
import "fmt"

func main()  {
	var no_produk, banyakBelanja int

	fmt.Scan(&no_produk, &banyakBelanja)
	if no_produk == 1 {
		fmt.Println(2980 * banyakBelanja)
	} else if no_produk == 2 {
		fmt.Println(4500 * banyakBelanja)
	} else if no_produk == 3 {
		fmt.Println(9980 * banyakBelanja)
	} else if no_produk == 4 {
		fmt.Println(4490 * banyakBelanja)
	} else if no_produk == 5 {
		fmt.Println(6870 * banyakBelanja)
	}
	
}