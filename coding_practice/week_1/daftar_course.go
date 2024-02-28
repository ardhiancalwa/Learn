package main

import "fmt"

func main() {
	var mat, fis, kim int
	var result bool

	fmt.Scan(&mat, &fis, &kim)
	result = (mat >= 65 && fis >= 55 && kim >= 50) && (mat+fis+kim == 180)
	fmt.Println(result)
}
