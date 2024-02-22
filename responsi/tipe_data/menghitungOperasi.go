package main
import "fmt"

func main()  {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	plusX := a*d + b*c
	plusY := b * d
	fmt.Println("penjumlahan:", plusX, "/",plusY)

	minX := a*d - b*c
	minY := b * d
	fmt.Println("pengurangan:", minX, "/",minY)

	timesX := a * c
	timesY := b * d
	fmt.Println("perkalian:", timesX, "/",timesY)

	divX := a * d
	divY := b * c
	fmt.Println("pembagian:", divX, "/",divY)
	
}