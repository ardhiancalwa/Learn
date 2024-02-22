package main
import "fmt"

func main()  {
	var bil, d1, d4 int

	fmt.Scan(&bil)
	d1 = bil / 1000
	d4 = bil % 10
	if d1%2 != 0 && d4%2 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}