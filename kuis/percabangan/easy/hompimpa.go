package main
import "fmt"

func main()  {
	var pemain1, pemain2, pemain3 string
	var pemenang string

	fmt.Scan(&pemain1, &pemain2, &pemain3)
	if pemain1 == pemain2 && pemain3 != pemain1 && pemain3 != pemain2 {
		pemenang = "pemain 3 pemenang"
	} else if pemain1 == pemain3 && pemain2 != pemain1 && pemain2 != pemain3 {
		pemenang = "pemain 2 pemenang"
	} else if pemain2 == pemain3 && pemain1 != pemain2 && pemain1 != pemain3 {
		pemenang = "pemain 1 pemenang"
	} else if pemain1 == pemain2 && pemain2 == pemain3 && pemain1 == pemain3 {
		pemenang = "imbang"
	}
	fmt.Println(pemenang)
}