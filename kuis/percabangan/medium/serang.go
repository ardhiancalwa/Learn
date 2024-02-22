package main
import "fmt"

func main()  {
	var sandi, kodeRahasia string

	for i := 1; i <= 6; i++ {
		fmt.Scan(&sandi)
		kodeRahasia += string(sandi[0])
	}

	if kodeRahasia == "serang" {
		fmt.Println("serang")
	} else {
		fmt.Println("tidak serang")
	}
}