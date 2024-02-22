package main
import "fmt"

func main()  {
	var total int
	total = 0
	for i := 2; i < 1000000001; i+=3 {
		total += i
	}
	fmt.Println(total)
}