package main
import "fmt"

func main()  {
	var i, n, m, jumlah int

	fmt.Scan(&n, &m)
	jumlah = 0
	i = n
	for {
		if i > m {
			break	
		}
		jumlah += i
		i++
	}
	fmt.Println(jumlah)
}
