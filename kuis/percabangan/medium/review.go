package main
import "fmt"

func main()  {
	var review1, review2, review3, review4 int
	const bagus float64 = 3.50
	const kurang float64 = 1.50
	fmt.Scan(&review1, &review2, &review3, &review4)
	if  float64(review1 + review2 + review3 + review4) / 4.0  >= bagus{
		fmt.Println("bagus")
	} else if float64(review1 + review2 + review3 + review4) / 4.0 <= kurang{
		fmt.Println("kurang")
	} else {
		fmt.Println("sedang")
	}

}