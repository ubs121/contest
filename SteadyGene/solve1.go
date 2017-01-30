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
	left := make([]GeneCount, n)
	right := make([]GeneCount, n)

	gi := map[byte]int{'G': 0, 'A': 1, 'C': 2, 'T': 3}

	// left chop counts
	left[0][gi[s[0]]]++
	l := 1
	for l < n && left[l-1][gi[s[l]]] < limit {
		// copy previous values
		for j := 0; j < 4; j++ {
			left[l][j] = left[l-1][j]
		}
		// increase counter
		left[l][gi[s[l]]]++
		l++
	}

	// right chop counts
	right[n-1][gi[s[n-1]]]++
	r := n - 2
	for r > 0 && right[r+1][gi[s[r]]] < limit {
		// copy previous values
		for j := 0; j < 4; j++ {
			right[r][j] = right[r+1][j]
		}
		// increase counter
		right[r][gi[s[r]]]++
		r--
	}

	// check combinations of each left[ll] and right[rr]
	min := n
	for ll := l - 1; ll >= 0; ll-- {
		r_ := r
		if r < ll {
			r_ = ll
		} else {
			// no more further check needed, because 'r_-ll-1' will increase
			if r_-ll > min {
				break
			}
		}

		for rr := r_ + 1; rr < n; rr++ {

			j := 0
			for j < 4 && left[ll][j]+right[rr][j] <= limit {
				j++
			}

			if j >= 4 {
				fmt.Println(ll, rr, left[ll], right[rr])

				if rr-ll-1 < min {
					min = rr - ll - 1
				}

				break
			}
		}

	}

	fmt.Println(min)

}
