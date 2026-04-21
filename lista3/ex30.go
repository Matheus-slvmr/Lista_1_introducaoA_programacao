package main

import (
	f "fmt"
	m "math"
)

func main() {
	var r, r3, v float64

	for r = 0; r <= 20; r += 0.5 {
		r3 = m.Pow(r, 3)
		v = 4 * m.Phi * r3 / 3
		f.Printf("Volume com R=%f: V=%fcm^3", r, v)
		f.Println()
	}
}