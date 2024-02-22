package main

import "fmt"

func main() {
	var N, M, i, fpb, kpk int

	fmt.Scan(&N, &M)
	for i = 1; i <= N; i++ {
		if N%i == 0 && M%i == 0 {
			fpb = i 
		}
	}
	kpk = (N*M) / fpb
	fmt.Println(kpk)
}
