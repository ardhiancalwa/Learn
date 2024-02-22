package main

import "fmt"

func main() {
	var bilBul, digitBefore, jumDigit, naik, turun int
	var bil float64

	fmt.Scan(&bil)
	bilBul = int(bil)
	digitBefore = bilBul%10
	bilBul /= 10
	for bilBul > 0 {
		jumDigit++
		if bilBul%10 < digitBefore {
			naik++
		} else if bilBul%10 > digitBefore {
			turun++
		}
		digitBefore = bilBul%10
		bilBul /= 10
	}
	if naik == jumDigit {
		fmt.Println("ascending")
	} else if turun == jumDigit {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}

	fmt.Println(naik, turun)
}