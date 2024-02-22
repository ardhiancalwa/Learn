package main

import "fmt"

func main() {
	var i, jumlah, digit1, digit2, n, bilangan int

	fmt.Scan(&n)
	jumlah = 0
	i = 1
	for i <= n{
		fmt.Scan(&bilangan)
		digit1 = bilangan / 1000 % 10 
		digit2 = bilangan % 100 / 10
		jumlah += digit1 + digit2
		i++
	}
	
	fmt.Println(jumlah)
}
