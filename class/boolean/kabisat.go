package main
import "fmt"

func main() {
	var kabisat bool	
	var tahun int

	fmt.Scan(&tahun)
	kabisat = (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0)
	fmt.Println(kabisat)
}