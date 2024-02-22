package main
import "fmt"

func main()  {
	var i, n, m int

	fmt.Scan(&n, &m)
	i = n
	for i <= m {
		fmt.Print(i, " ")
		i++
	}
	
}