package main

import "fmt"

func main() {
	var a1, razao, n int
	fmt.Scan(&a1, &razao, &n)

	soma := 0
	for i := 0; i < n; i++ {
		pa := a1 + i*razao
		soma += pa
	}
	fmt.Printf("%d\n", soma)
}