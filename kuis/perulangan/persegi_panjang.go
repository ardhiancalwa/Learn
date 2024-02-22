package main
import "fmt"

func main()  {
	var i, n, luas, keliling, panjang, lebar int
	
	fmt.Scan(&n)
	i = 1
	for i <= n {
		fmt.Scan(&panjang, &lebar)
		luas = panjang * lebar
		keliling = 2 * (panjang + lebar)
		fmt.Println(luas, keliling)
		i++
	}
}