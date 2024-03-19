package main

import "fmt"

func main() {
	var n int
	var result string
	fmt.Scan(&n)
	result = polaBintang(n, 1, 0, "")
	fmt.Print(result)
}

func polaBintang(n, i, j int, result string) string {
	if i > n {
		return result
	}
	if i > j {
		return polaBintang(n, i, j+1, result+"*")
	} else {
		return polaBintang(n, i+1, 0, result+"\n")
	}
}
