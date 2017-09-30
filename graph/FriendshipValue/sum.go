package main

import "fmt"

func main() {
	n := 100

	for i := 0; i <= n; i++ {
		j := n - i
		si := ((i-1)*i*(i+1))/3 + j*(j-1)
		fmt.Println(i, j, si)

	}
}
