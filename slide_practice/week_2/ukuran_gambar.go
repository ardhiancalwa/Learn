package main

import "fmt"

func zoomIn(p, l int, skala int) {
	var w, h int
	if skala > 0 && skala <= 100 {
		w = p * skala
		h = l * skala
		fmt.Println(w, h)
	}

}
func zoomOut(p, l, skala int) {
	var w, h int
	if skala > 0 && skala <= 100 {
		w = p / skala
		h = l / skala
		fmt.Println(w, h)
	}

}

func main() {
	var p, l, skala int
	var s string
	fmt.Scan(&p, &l)
	fmt.Scan(&s, &skala)
	if s == "+" {
		zoomIn(p, l, skala)
	} else if s == "-" {
		zoomOut(p, l, skala)
	}
}
