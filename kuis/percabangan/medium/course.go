package main
import "fmt"

func main()  {
	var mat, fis, kim int
	var daftar bool

	fmt.Scan(&mat, &fis, &kim)
	daftar = (mat >= 65 && fis >= 55 && kim >= 50) && (mat + fis + kim >= 180)
	fmt.Println(daftar)
}