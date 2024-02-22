package main
import "fmt"

func main()  {
	var i, n, luas, keliling,  sisi int

	fmt.Scan(&n)
	i = 1
	for i <= n {
		fmt.Scan(&sisi)
		luas = sisi * sisi
		keliling = sisi * 4
		i++
		fmt.Println(luas, keliling)
	}
}