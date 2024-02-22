package main
import "fmt"

func main()  {
	var hariKerja, genreAksi bool
	var jamNonton int
	var jadiNonton bool

	fmt.Scan(&hariKerja, &jamNonton, &genreAksi)
	if hariKerja && jamNonton > 19  && genreAksi {
		jadiNonton = true
	} else {
		jadiNonton = false
	}
	fmt.Println(jadiNonton)
}