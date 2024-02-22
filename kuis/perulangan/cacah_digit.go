package main
import "fmt"

func main()  {
	var digit, bilangan int

	fmt.Scan(&bilangan)
	for bilangan > 0 {
		digit = bilangan % 10
		fmt.Print(digit, " ")
		bilangan /= 10
	}
}