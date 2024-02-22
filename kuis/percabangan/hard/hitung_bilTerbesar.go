package main

import "fmt"

func main() {
	var bil1, bil2, bil3, bil4 int
	// var max int

	fmt.Scan(&bil1, &bil2, &bil3, &bil4)
	if bil1 >= bil2 && bil1 >= bil3 && bil1 >= bil4 {
		fmt.Println(bil1)
	} else if bil2 >= bil1 && bil2 >= bil3 && bil2 >= bil4 {
		fmt.Println(bil2)
	} else if bil3 >= bil1 && bil3 >= bil2 && bil3 >= bil4 {
		fmt.Println(bil3)
	} else if bil4 >= bil1 && bil4 >= bil2 && bil3 >= bil3 {
		fmt.Println(bil4)
	}
}
