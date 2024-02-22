package main
import "fmt"

func main(){
	var i, n, faktor int
	
	faktor = 1
	fmt.Scan(&n)
	for i = 5; i >= n; i-= 1{
		faktor *= i
	}
	fmt.Println(faktor)
}