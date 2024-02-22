package main
import "fmt"

func main()  {
	var i, n int
	var kata_kata string

	fmt.Scan(&n, &kata_kata)
	i = 1
	for i <= n {
		fmt.Println(kata_kata)
		i++
	}
	fmt.Println(i) 
	
}