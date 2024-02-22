package main
import "fmt"

func main()  {
	var n, i int
	var i1, i2, i3, i4, i5, siap bool

	fmt.Scan(&n)
	i = 1
	siap = true
	for i <= n && siap {
		fmt.Scan(&i1, &i2, &i3, &i4, &i5)
		siap = i1 && i2 && i3 && i4 && i5
		i++
	}
	fmt.Println(siap)
}