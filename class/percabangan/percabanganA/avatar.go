package main
import "fmt"

func main()  {
	var P, dewasa, kecil, sisa int

	fmt.Scan(&P)
	if P <= 15 {
		dewasa = P / 5
		if P % 5 > 0 {
			dewasa += 1
		}
		fmt.Println("Dewasa", dewasa)
	} else if P > 15 {
		dewasa = 3
		sisa = P - dewasa * 5
		if sisa <= 10 {
			kecil = sisa / 2
			if sisa % 2 > 0 {
				kecil += 1
			}
			fmt.Println("Dewasa", dewasa, "Kecil", kecil)
		} else if sisa > 10 {
			kecil = 5
			sisa = sisa - kecil * 2
			fmt.Println("Dewasa", dewasa, "Kecil", kecil, "dan", sisa, "tidak berangkat")
		}
	} 
}