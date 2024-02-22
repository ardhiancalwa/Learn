package main
import "fmt"

func main()  {
	var i, n, m, jumlah int

	fmt.Scan(&n, &m)
	jumlah = 0
	i = n
	for i <= m {
		jumlah += i
		i++
	}
	fmt.Println(jumlah)
}