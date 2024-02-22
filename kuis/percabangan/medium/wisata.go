package main
import "fmt"

func main()  {
	var k, p int

	fmt.Scan(&k, &p)
	if p % k != 0 {
		fmt.Println((p/k) + 1)
	} else {
		fmt.Println(p/k)
	}
}