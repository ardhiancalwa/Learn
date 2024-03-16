package main

import "fmt"

func main() {
	var N int
	var result string
	fmt.Scan(&N)
	result = printStars(N, 1, 0, "")
	fmt.Print(result)
}

func printStars(N, i, j int, result string) string {
	if i > N {
		return result
	}
	if j < i {
		return printStars(N, i, j+1, result+"*")
	} else {
		return printStars(N, i+1, 0, result+"\n")
	}
}
