package main
import "fmt"

func main()  {
	var bilangan, digit int

	fmt.Scan(&bilangan)
	digit = 0
	for bilangan != 0 {
		bilangan /= 10
		digit++
	}
	fmt.Println(digit)
}