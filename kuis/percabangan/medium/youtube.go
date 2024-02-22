package main
import "fmt"

func main()  {
	var pelanggan, jam int

	fmt.Scan(&pelanggan, &jam)
	if pelanggan >= 1000 && jam >= 4000 {
		fmt.Println("sudah dapat dimonetisasi")
	} else {
		fmt.Println("belum dapat dimonetisasi")
	}
}