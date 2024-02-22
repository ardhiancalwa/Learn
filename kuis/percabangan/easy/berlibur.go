package main
import "fmt"

func main()  {
	var gaji, bonus bool
	var berlibur string

	fmt.Scan(&gaji, &bonus)
	if gaji && bonus {
		berlibur = "berlibur"
	} else if gaji || bonus {
		berlibur = "berlibur"
	} else {
		berlibur = "tidak berlibur"
	}
	fmt.Println(berlibur)
}