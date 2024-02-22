package main
import "fmt"

func main()  {
	var umur int
	var fasih_golang, ktp bool

	fmt.Scan(&umur, &fasih_golang)
	if umur >= 10 && fasih_golang {
		ktp = true
	} else {
		ktp = false
	}
	fmt.Println(ktp)
}