package main

import "fmt"

func main() {
	var bil, max, secondMax int

	fmt.Scan(&bil)
	max = bil
	secondMax = bil

	for bil != 0 {
		fmt.Scan(&bil)
		if bil > max {
			secondMax = max
			max = bil
		} else if bil > secondMax && bil != max {
			secondMax = bil
		}
	}
	fmt.Println(secondMax)
}
