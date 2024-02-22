package main
import "fmt"

func main()  {
	var hujan, kondisi string

	fmt.Scan(&hujan)
	if hujan == "tinggi" {
		kondisi = "macet"
	} else if hujan == "rendah" {
		kondisi = "tidak macet"
	}
	fmt.Println(kondisi)
}