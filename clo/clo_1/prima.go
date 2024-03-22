package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))

}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}