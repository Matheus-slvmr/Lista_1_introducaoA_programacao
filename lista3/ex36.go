package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var n, n1 int
	var base string

	f.Println("Insira um número na base decimal :")
	f.Scan(&n)
	n1 = n
	if n == 0 {
		base = " 00"
	}
	for n > 0 {
		resto := n % 16
		if resto == 0 {
			base = " 00" + base
		} else {
			if resto < 10 {
				base = strconv.Itoa(resto) + base
			} else {
				if resto == 10 {
					base = "A" + base
				}
				if resto == 11 {
					base = "B" + base
				}
				if resto == 12 {
					base = "C" + base
				}
				if resto == 13 {
					base = "D" + base
				}
				if resto == 14 {
					base = "E" + base
				}
				if resto == 15 {
					base = "F" + base
				}
			}
		}
		n = n / 16
	}
	f.Printf(" %d na base 16 é: %s", n1, base)
}