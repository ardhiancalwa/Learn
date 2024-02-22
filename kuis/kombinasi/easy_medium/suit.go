package main

import "fmt"

func main() {
	var i, n int
	var A, B int
	var pemain1, pemain2 byte

	fmt.Scan(&n)
	
	for i = 1; i <= n; i++ {
		fmt.Scanf("%c%c", &pemain1, &pemain2)
		if (pemain1 == 'g' && pemain2 == 'k') || (pemain1 == 'b' && pemain2 == 'g') || (pemain1 == 'k' && pemain2 == 'b') {
			A++
		} else if (pemain1 == 'k' && pemain2 == 'g') || (pemain1 == 'g' && pemain2 == 'b') || (pemain1 == 'b' && pemain2 == 'k'){
			B++
		}
	}
	fmt.Println(A, B)
}