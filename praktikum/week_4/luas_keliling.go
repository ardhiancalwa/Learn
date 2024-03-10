package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {

	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, toLuas, toKel *float64) {
	*toLuas = lP + lL
	*toKel = kL + kP
}

const pi = 3.14

func main() {
	var radius, sisi float64
	var l, k, lL, lP, kL, kP, toLuas, toKel float64
	fmt.Scan(&radius, &sisi)

	if radius != 0 && sisi != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
		for radius != 0 && sisi != 0 {
			lL = pi * radius * radius
			lP = sisi * sisi
			kL = 2 * pi * radius
			kP = 4 * sisi
			hitungLuasKelilingLingkaran(radius, &l, &k)
			hitungLuasKelilingPersegi(sisi, &l, &k)
			hitungTotal(lL, lP, kL, kP, &toLuas, &toKel)
			fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", radius, sisi, lL, lP, kL, kP, toLuas, toKel)
			fmt.Scan(&radius, &sisi)
		}
	}
}
