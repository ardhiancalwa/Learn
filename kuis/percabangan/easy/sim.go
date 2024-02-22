package main
import "fmt"

func main()  {
	var umur, tesMengemudi int
	var fasih_golang, SIM bool

	fmt.Scan(&umur, &fasih_golang, &tesMengemudi)
	if umur >= 10 && tesMengemudi == 100 && fasih_golang {
		SIM = true
	} else {
		SIM = false
	}
	fmt.Println(SIM)
}