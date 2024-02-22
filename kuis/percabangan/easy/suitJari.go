package main
import "fmt"

func main()  {
	var pemain1, pemain2 string
	var pemenang string

	fmt.Scan(&pemain1, &pemain2)
	if pemain1 == "jempol" && pemain2 == "telunjuk" {
		pemenang = "pemain 1 menang"
	} else if pemain1 == "jempol" && pemain2 == "kelingking" {
		pemenang = "pemenang 2 menang"
	} else if pemain1 == "telunjuk" && pemain2 == "jempol" {
		pemenang = "pemain 2 menang"
	} else if pemain1 == "telunjuk" && pemain2 == "kelingking" {
		pemenang = "pemain 1 menang"
	} else if pemain1 == "kelingking" && pemain2 == "jempol" {
		pemenang = "pemain 1 menang"
	} else if pemain1 == "kelingking" && pemain2 == "telunjuk" {
		pemenang = "pemain 2 menang"
	} else if pemain1 == pemain2 {
		pemenang = "imbang"
	}
	fmt.Println(pemenang)
}