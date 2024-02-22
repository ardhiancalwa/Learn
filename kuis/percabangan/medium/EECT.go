package main
import "fmt"

func main()  {
	var P, V, GA, FS, CI, TC int
	var rata float64

	
	fmt.Scan(&P, &V, &GA, &FS, &CI, &TC)
	rata = float64(P+V+GA+FS+CI+TC) / 6
	if  rata >= 3.33 {
		fmt.Println("mendapat sertifikat")
	} else {
		fmt.Println("tidak mendapat sertifikat")
	}
}