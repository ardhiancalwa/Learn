package main
import "fmt"

func main()  {
	var jumGol, gol, i int
	var hadiah bool
	fmt.Scan(&gol)
	jumGol = 0
	for i = 1; gol >= 10 || gol != -1; i++ {
		jumGol += gol 
		fmt.Scan(&gol)
	}
	hadiah = jumGol >= 10
	fmt.Println(hadiah)
}	