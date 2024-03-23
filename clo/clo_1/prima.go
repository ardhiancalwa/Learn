package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))

}

func isPrime(n int) bool {
	var jum int
	for i := 1; i <= n; i++ {
        if n%i == 0 {
            jum++
        }
    }
    return jum == 2
}
