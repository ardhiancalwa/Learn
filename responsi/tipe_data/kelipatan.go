package main
import "fmt"

func main()  {
	var bil int
	var check bool

	fmt.Scan(&bil)
	check = (bil % 3 == 0) && (bil % 5 == 0)
	fmt.Println(check)
}