package main

import "fmt"

func main() {
	var jenis string
	var jumPelajar int
	var sd, smp, sma int

	fmt.Scan(&jenis, &jumPelajar)
	for jenis != "x" && jumPelajar != 0 {
		if jenis == "sd" {
            sd += jumPelajar
        }
        if jenis == "smp" {
            smp += jumPelajar
        }
        if jenis == "sma" {
            sma += jumPelajar
        }
        fmt.Scan(&jenis, &jumPelajar)
	}
	fmt.Println(sd, smp, sma)
}