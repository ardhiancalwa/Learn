package main

import "fmt"

func hitungmenang(g, k int, jm *int) {
	if g > k {
		*jm = *jm + 1
	}
}
func hitungdraw(g, k int, jd *int) {
	if g == k {
		*jd = *jd + 1
	}
}
func hitungkalah(g, k int, jk *int) {
	if g < k {
		*jk = *jk + 1
	}
}
func hitungjumgolkegolanselisih(g, k int, jg, jk, jsg *int) {
	*jg = *jg + g
	*jk = *jk + k
	*jsg = *jg - *jk
}
func hitungjumpoint(jp *int) {

	*jp = hitungmenang(g, k, &jm)*3 + hitungdraw(g, k, &jd)
}

func main() {
	var n, g, k, i, jm, jd, jk, jk1, jg, jsg, jp int
	fmt.Scan(&n)
	i = 1
	for i <= n {
		fmt.Scan(&g, &k)
		hitungmenang(g, k, &jm)
		hitungdraw(g, k, &jd)
		hitungkalah(g, k, &jk)
		hitungjumgolkegolanselisih(g, k, &jg, &jk1, &jsg)
		hitungjumpoint(&jp)
		i = i + 1
	}

	fmt.Print(n, " ", jm, " ", jd, " ", jk, " ", jg, " ", jk, " ", jsg, " ", jp)

}