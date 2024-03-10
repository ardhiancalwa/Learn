package main

import "fmt"

func hitungmenang(g, k int, jm *int) {
	if g > k {
		*jm = *jm + 1
	}
}
func hitungdraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}
func hitungkalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}
func hitungjumgolkegolanselisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg = *jg - *jk
}
func hitungjumpoint(g, k int, jm, jd, jp *int) {
	if g > k {
		*jm += 1
	}
	if g == k {
		*jd += 1
	}
	*jp = (*jm*3) + (*jd*1)
}

func main() {
	var n, g, k, i, jm, jd, jk, jk1, jg, jsg, jp int
	fmt.Scan(&n)
	
	for i = 1; i <= n; i++ {
		fmt.Scan(&g, &k)
		hitungkalah(g, k, &jk)
		hitungjumgolkegolanselisih(g, k, &jg, &jk1, &jsg)
		hitungjumpoint(g, k, &jm, &jd, &jp)	
	}
	fmt.Print(n, " ", jm, " ", jd, " ", jk, " ", jg, " ", jk1, " ", jsg, " ", jp)
}
