package main
import "fmt"

func main()  {
	var h1, m1, h2, m2 int
	var hh, mm int

	fmt.Scan(&h1, &m1, &h2, &m2)
	if h2 < h1 {
		h2 += 12
	}

	h1 *= 60
	h2 *= 60

	hh = (h2 - h1) / 60
	mm = m2 - m1

	fmt.Println(hh, "jam", mm, "menit")
}