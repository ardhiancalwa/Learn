package main
import "fmt"

func main()  {
	var i, x, y, n, jumlah int

	fmt.Scan(&x, &y, &n)
	jumlah = 0
	i = x
	for i <= n {
		jumlah += 1
		i += y + x
	}

	i = x + y
	for i <= n {
		jumlah += 1
		i += x + y
	}
	fmt.Println(jumlah)
}