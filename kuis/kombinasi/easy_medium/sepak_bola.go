package main

import "fmt"

func main() {
	var jumPoin, jumGol, jumKemasukan, selisih int
	var gol, kemasukan int
	var i int
	const match int = 38

	jumPoin = 0
	jumGol = 0
	jumKemasukan = 0
	for i = 1; i <= match; i++ {
		fmt.Scan(&gol, &kemasukan)
		jumGol += gol
		jumKemasukan += kemasukan
		if gol > kemasukan {
			jumPoin += 3
		} else if gol == kemasukan {
			jumPoin += 1
		} else {
            jumPoin += 0
        }
	}
	selisih = jumGol - jumKemasukan
	fmt.Println(jumPoin, jumGol, jumKemasukan, selisih)
}
