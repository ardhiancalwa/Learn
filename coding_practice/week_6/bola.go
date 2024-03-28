package main

import "fmt"

func main() {
	type dataPemain struct {
		nama   string
		gol    int
		assist int
	}
	var rataGol, rataAssist  float64
	var jumGol, jumAssist int
	var bola dataPemain
	for i := 1; i <= 3; i++ {
		fmt.Scan(&bola.nama, &bola.gol, &bola.assist)
		jumGol += bola.gol
		jumAssist += bola.assist
	}
	rataGol = float64(jumGol) / float64(3)
	rataAssist = float64(jumAssist) / float64(3)
	fmt.Println(rataGol)
	fmt.Println(rataAssist)
}
