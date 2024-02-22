package main
import "fmt"

func main()  {
	var hujan, mood string
	var kondisi string

	fmt.Scan(&hujan, &mood)
	if hujan == "true" {
		if mood == "true" {
			kondisi = "keluar jalan-jalan"
		} else {
			kondisi = "diam di rumah saja"
		}
	} else {
		kondisi = "diam di rumah saja"
	}
	fmt.Println(kondisi)
}