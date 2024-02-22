package main
import "fmt"

func main()  {
	var bil1, bil2, selisih int

	fmt.Scan(&bil1, &bil2)
	selisih = bil1 - bil2
	if selisih < 0 {
		selisih *= -1
	}
	fmt.Println("Selisih angka Heiji dan Shinichi adalah", selisih)
}