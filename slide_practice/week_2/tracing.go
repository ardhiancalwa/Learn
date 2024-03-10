package main

import "fmt"

func procB(b, c *int) {
	*b += *c
	*c += a + *b
}

var a int

func main() {
	a = 10
	procB(&a, &a)
	fmt.Println(a)
}
