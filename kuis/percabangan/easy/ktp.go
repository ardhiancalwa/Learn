package main
import "fmt"

func main()  {
	var umur int
	var kk bool
	var ktp string

	fmt.Scan(&umur, &kk)
	if umur >= 17 && kk {
		ktp = "bisa membuat KTP"
	} else {
		ktp = "belum bisa membuat KTP"
	}
	fmt.Println(ktp)
}