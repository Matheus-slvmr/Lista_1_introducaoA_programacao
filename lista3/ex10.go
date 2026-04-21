package main

import f "fmt"
func main (){
	
	var n int
	f.Scan(&n)
	if n <= 3 {
		return
	}
	a := 0
	b := 1
	f.Printf("%d - %d", a, b)
		for i := 3; i <= n;i++{
			proximo := a + b 
			f.Printf(" - %d", proximo)
			a=b
			b=proximo
			
		}
	
}