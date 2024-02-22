package main

import "fmt"

func main() {
	var episode, durasi int
	var jumDurasi int

	fmt.Scan(&episode, &durasi)
	for episode != 0 && durasi != 0 {
		if episode%2 != 0 {
			jumDurasi += durasi
		}
		fmt.Scan(&episode, &durasi)
	}
	fmt.Println(jumDurasi)
}