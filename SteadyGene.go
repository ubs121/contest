// https://www.hackerrank.com/challenges/bear-and-steady-gene

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var s string
	fmt.Scanf("%s", &s)

	//limit := n / 4

	m := map[byte]int{'A': 0, 'C': 0, 'T': 0, 'G': 0}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	fmt.Println(m)
}

// GAAATAAA
// GAAGTTCC

//  6 1 1 0
// -4 +1 +1 +2
