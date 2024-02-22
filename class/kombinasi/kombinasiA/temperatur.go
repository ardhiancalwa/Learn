package main

import "fmt"

func main() {
	var suhu, stop, n, rata, max, min int
	var jumRata float64
	stop = 1
	n = 0
	max = 0
	min = 0
	for suhu+stop != 0 {
		stop = suhu
		fmt.Scan(&suhu)
		n++
		rata += suhu
		if suhu > max {
			max = suhu
		}
		if suhu < min {
			min = suhu
		}
	}
	jumRata = float64(rata) / float64(n-1)
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(jumRata)
}