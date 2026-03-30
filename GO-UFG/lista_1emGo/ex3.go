package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)
	nf := ((n1 * 100) + (n2 * 10) + n3)
	nf2 := nf * nf
	fmt.Println(nf, nf2)
}