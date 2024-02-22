package main

import "fmt"

func main() {
	var b1, b2, b3, b4, max, min int
	var mean, median float64

	fmt.Scan(&b1, &b2, &b3, &b4)
	mean = float64(b1+b2+b3+b4) / 4
	max = b1
	if b2 >= max {
		max = b2
	} else {
		max = max
	}
	if b3 >= max {
		max = b3
	} else {
		max = max
	}
	if b4 >= max {
		max = b4
	} else {
		max = max
	}

	min = b1
	if b2 <= min {
		min = b2
	} else {
		min = min
	}
	if b3 <= min {
		min = b3
	} else {
		min = min
	}
	if b4 <= min {
		min = b4
	} else {
		min = min
	}
	fmt.Println(min, max)

	if b2 >= b3 {
		fmt.Println(min, b3, b2, max)
		median = float64(b3+b2) / 2
		fmt.Println(mean, median)
		if mean == median {
			fmt.Println("simetris")
		} else if mean >= median {
			fmt.Println("skew kanan")
		} else if mean <= median {
			fmt.Println("skew kiri")
		}
	} else if b3 >= b2 {
		fmt.Println(min, b2, b3, max)
		median = float64(b2+b3) / 2
		fmt.Println(mean, median)
		if mean == median {
			fmt.Println("simetris")
		} else if mean >= median {
			fmt.Println("skew kanan")
		} else if mean <= median {
			fmt.Println("skew kiri")
		}
	}

}
