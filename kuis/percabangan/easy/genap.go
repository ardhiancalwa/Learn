package main
import "fmt"

func main() {
	var bil int
	var genap string

	fmt.Scan(&bil)
	if bil%2 == 0 && bil >= 0 {
		genap = "ya"
	} else if bil%2 == 0 && bil < 0 {
		genap = "tidak"
	} else {
		genap = "tidak"
	}
	fmt.Println(genap)
}
