package main

import (
	f "fmt"
	m "math"
)

func main() {

	s := 0.0

	for i := 0; i <= 50; i++ {
		s = s + m.Pow(-1, float64(i))/m.Pow((1+2*float64(i)), 3)
	}
	k := s * 32
	y := m.Pow(k, 1.0/3.0)

	f.Println("o valor de PI é =", y)
}