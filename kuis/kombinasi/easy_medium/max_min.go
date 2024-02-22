package main

import "fmt"

func main() {
	var max, min int
	var bil int

	fmt.Scan(&bil)
	max = bil
	min = bil
	for bil != 0 {
		if bil > max {
			max = bil
		}
		if bil < min {
            min = bil
        }
		fmt.Scan(&bil)
	}
	fmt.Println(max, min)
}