package main
import "fmt"

func main()  {
	var daerah1, daerah2 byte
	var tetangga bool

	fmt.Scanf("%c %c", &daerah1, &daerah2)
	if daerah1 == 'A' && daerah2 == 'B' || daerah1 == 'B' && daerah2 == 'A' {
		tetangga = true
	} else if daerah1 == 'B' && daerah2 == 'D' || daerah1 == 'D' && daerah2 == 'B' {
		tetangga = true
	} else if daerah1 == 'B' && daerah2 == 'C' || daerah1 == 'C' && daerah2 == 'B' {
		tetangga = true
	} else if daerah1 == 'C' && daerah2 == 'E' || daerah1 == 'E' && daerah2 == 'C' {
		tetangga = true
	} else {
		tetangga = false
	}
	fmt.Println(tetangga)
}