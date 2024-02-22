package main
import "fmt"

func main()  {
	var N, d int
	var s bool

	fmt.Scan(&N)
	d = 1
	for d <= N {
		fmt.Scan(&d)
		s = N % d == 0
		fmt.Println(d, s)
		d++
	}
}