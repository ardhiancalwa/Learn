package main
import "fmt"

func main()  {
	var shampoo, hujan string
	var kondisi string

	fmt.Scan(&shampoo, &hujan)
	if shampoo == "ya" {
		if hujan == "ya" {
			kondisi = "tidak pergi ke minimarket"
		} else {
			kondisi = "pergi ke minimarket"
		}
	} else {
		kondisi = "tidak pergi ke minimarket"
	}
	fmt.Println(kondisi)
}