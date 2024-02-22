package main
import "fmt"

func main()  {
	var i, n int
	var kata_kata string	

	fmt.Scan(&n)
	fmt.Scan(&kata_kata)
	for i = 1; i <= n; i++ {
		fmt.Println(kata_kata)
	}
}