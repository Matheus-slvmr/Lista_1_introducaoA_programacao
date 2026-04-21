package main

import (
	f "fmt"
	m "math"
)

func main() {
	s := 0.0

	for i := 0; i <= 14; i++ {
		s = s + m.Pow(-2, float64(i))/(m.Pow(15-float64(i), 2))
	}
	f.Println("S =", s)
}