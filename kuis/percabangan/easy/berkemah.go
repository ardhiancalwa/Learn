package main
import "fmt"

func main()  {
	var sehat, bekal, cuaca bool
	var kemah string

	fmt.Scan(&sehat, &bekal, &cuaca)
	if sehat && bekal && cuaca {
		kemah = "berkemah"
	} else {
		kemah = "tidak berkemah"
	}
	fmt.Println(kemah)
}