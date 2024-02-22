package main
import "fmt"

func main()  {
	var bilangan1, bilangan2, d1, d2, n int 
	fmt.Scan(&bilangan1)
	d1 = bilangan1 / 10
	d2 = bilangan1 % 10
	n = 1
	for {
		fmt.Scan(&bilangan2)
		if bilangan1 >= 10 && bilangan1 <= 99 && bilangan2 == d1 || bilangan2 == d2 || bilangan2 == bilangan1{
			break	
		}
		n++
	}
	fmt.Println(n)
}