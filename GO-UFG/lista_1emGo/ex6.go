package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var fahrenheitr []float64
	var celciusr []float64
	for i := 0; i < n; i++ {
		var fahrenheit float64
		fmt.Scan(&fahrenheit)

		celcius := (fahrenheit - 32) * 5 / 9
		fahrenheitr = append(fahrenheitr, fahrenheit)
		celciusr = append(celciusr, celcius)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%.2f\n FAHRENHEIT EQUIVALE A %.2f\n CELCIUS ", fahrenheitr[i], celciusr[i])
	}
}