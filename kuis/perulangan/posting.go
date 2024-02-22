package main
import "fmt"

func main()  {
	var i, n, hari, postingan, text, gambar int

	fmt.Scan(&n)
	text = 0
	gambar = 0
	i = 1
	for i <= n {
		fmt.Scan(&hari)
		text += hari
		fmt.Scan(&postingan)
		gambar += postingan
		i++
	}
	fmt.Println(text, gambar)
}