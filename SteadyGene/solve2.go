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
	right := make([]GeneCount, n+1) // counter right[position][G,A,C,T]

	ind := map[byte]int{'G': 0, 'A': 1, 'C': 2, 'T': 3}

	right[n-1][ind[s[n-1]]]++
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1]
		right[i][ind[s[i]]]++
	}

	total := count[0]

	min := n - 1
	var left, middle GeneCount

	c := 0
	// check combinations of each left[i] and right[j]
	for i := 0; i < n && left[ind[s[i]]] < limit; i++ {
		left[ind[s[i]]]++ // increase counter

		for j := n - 1; j > i && right[j][ind[s[j]]] <= limit-left[ind[s[j]]]; j-- {
			c++
			if min > j-i-1 {
				min = j - i - 1
			}
		}
	}

	fmt.Println(min, c)

}
