package main

import "fmt"

func main() {
	var ada bool = false
	var hrf byte
	fmt.Scanf("%c", &hrf)
	for hrf != '.' && !ada{
		 ada = huruf(hrf) 
		 fmt.Scanf("%c", &hrf)
	} 
	fmt.Println(ada)
}

func huruf(hrf byte) bool {
	if hrf == 'k' || hrf == 'q' {
		return true
	}
	return false
}
