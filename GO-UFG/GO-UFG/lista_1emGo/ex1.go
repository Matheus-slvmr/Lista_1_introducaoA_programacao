package main

import "fmt"

func main()  {
	var(
		n1 , n2 , n3 float32
		media float32
	)
	fmt.Scan(&n1 , &n2 , &n3)
	media = (n1 + n2 + n3) / 3
	if media >= 6.0 {
		fmt.Println("sua media é:",media)
		fmt.Println("aprovado")
	}
	

}