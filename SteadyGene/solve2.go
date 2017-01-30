// https://www.hackerrank.com/challenges/bear-and-steady-gene

package main

import "fmt"

// 40
// TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC
// 5

type GeneCount [4]int

func main() {

	var n int // 4<=n<=500'000
	fmt.Scanf("%d", &n)

	var s string
	fmt.Scanf("%s", &s)

	limit := n / 4
	gc := make([]GeneCount, n) // counter gc[position][G,A,C,T]

	gi := map[byte]int{'G': 0, 'A': 1, 'C': 2, 'T': 3}

	gc[0][gi[s[0]]]++

	l := 1
	for l < n {
		// copy previous values
		for j := 0; j < 4; j++ {
			gc[l][j] = gc[l-1][j]
		}
		// increase counter
		gc[l][gi[s[l]]]++
		l++
	}

	fmt.Println(limit, gc)

}
