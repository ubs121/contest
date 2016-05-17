package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		for j := n - i; j < n; j++ {
			fmt.Print("#")
		}

		fmt.Println()
	}
}
