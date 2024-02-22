// package main

// import "fmt"

// func main() {
// 	var karakter byte
// 	var jumNumerik int
// 	fmt.Scanf("%c", &karakter)
// 	for karakter != '.' {
// 		if karakter >= '0' && karakter <= '9' {
// 			jumNumerik += int(karakter - '0')
// 		}
// 		fmt.Scanf("%c", &karakter)
// 	}

// 	fmt.Println(jumNumerik)
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var kalimat string
	var total, i int

	fmt.Scan(&kalimat)
	for i = 0; i < len(kalimat); i++ {
		var bil, err = strconv.Atoi(string(kalimat[i]))

		if err == nil {
			total += bil
		}
	}
	fmt.Println(total)
}