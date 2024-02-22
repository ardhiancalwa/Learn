package main
import "fmt"

func main() {
	var poin1, poin2, poin3, poin4, poin5 int

	fmt.Scan(&poin1, &poin2, &poin3, &poin4, &poin5)
	if poin1+poin2+poin3 == poin3+poin4+poin5 && (poin2+poin3+poin4)%(poin1+poin5) == 0 {
		fmt.Println("cashback")
		fmt.Println("diskon")
	} else if poin1+poin2+poin3 == poin3+poin4+poin5 {
		fmt.Println("cashback")
	} else if (poin2+poin3+poin4)%(poin1+poin5) == 0 {
		fmt.Println("diskon")
	} else {
		fmt.Println(" ")
	}
}
