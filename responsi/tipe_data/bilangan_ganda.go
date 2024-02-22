package main
import "fmt"

func main()  {
	var digit1, digit2, bil int

	fmt.Scan(&bil)

	digit1 = bil / 10
	digit2 = bil % 10
	bil = (digit1 * 1000 + digit1 * 100 + digit2 * 10 + digit2)
	fmt.Println(bil)
}