package main

import "fmt"

func main() {
	var jumGol1, jumGol2, jumGol3, jumGol4 int
	var terbanyak, tersedikit int

	fmt.Scan(&jumGol1, &jumGol2, &jumGol3, &jumGol4)

	if jumGol1 >= jumGol2 && jumGol1 >= jumGol3 && jumGol1 >= jumGol4 {
		terbanyak = jumGol1
	} else if jumGol2 >= jumGol1 && jumGol2 >= jumGol3 && jumGol2 >= jumGol4 {
		terbanyak = jumGol2
	} else if jumGol3 >= jumGol1 && jumGol3 >= jumGol2 && jumGol3 >= jumGol4 {
		terbanyak = jumGol3
	} else if jumGol4 >= jumGol1 && jumGol4 >= jumGol2 && jumGol4 >= jumGol3 {
		terbanyak = jumGol4
	}

	if jumGol1 <= jumGol2 && jumGol1 <= jumGol3 && jumGol1 <= jumGol4 {
		tersedikit = jumGol1
	} else if jumGol2 <= jumGol1 && jumGol2 <= jumGol3 && jumGol2 <= jumGol4 {
		tersedikit = jumGol2
	} else if jumGol3 <= jumGol1 && jumGol3 <= jumGol2 && jumGol3 <= jumGol4 {
		tersedikit = jumGol3
	} else if jumGol4 <= jumGol1 && jumGol4 <=jumGol2 && jumGol4 <= jumGol3 {
		tersedikit = jumGol4
	}
	fmt.Println(terbanyak, tersedikit)
}
