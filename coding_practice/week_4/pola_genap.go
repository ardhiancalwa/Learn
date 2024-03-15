package main

import "fmt"

func main() {
	var N int
	var result string
	fmt.Scan(&N)
	result = printStars(N, 1, 0, "")
	fmt.Print(result)
}

func printStars(N int, i int, j int, result string) string {
	if i == N {
		return result
	} else {
		result += "*"
		return printStars(N, i+1, j+1, result)
	}
}
