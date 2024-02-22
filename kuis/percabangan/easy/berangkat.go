package main
import "fmt"

func main()  {
	var hujan_turun, punya_payung string
	var kondisi string

	fmt.Scan(&hujan_turun, &punya_payung)
	if hujan_turun == "ya" {
		if punya_payung == "ya" {
			kondisi = "berangkat"
		} else if punya_payung == "tidak" {
			kondisi = "diam di rumah"
		}
	} else if hujan_turun == "tidak" {
		kondisi = "berangkat"
	}
	fmt.Println(kondisi)
	
}