package main
import "fmt"

func main()  {
	var bil int
	var mutlak int

	fmt.Scan(&bil)
	if 	bil >= 0 {
		mutlak = bil
	} else if bil <= 0 {
		mutlak = bil - bil
		mutlak -= bil
	}
	fmt.Println(mutlak)
}