// https://www.hackerrank.com/challenges/bear-and-steady-gene

package main

import "fmt"
import "os"
import "bufio"

// 40
// TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC
// 5

type GeneCount [4]int

func main() {
	file, _ := os.Open("input10.txt")
	defer file.Close()

	buf := make([]byte, 500100)
	scanner := bufio.NewScanner(file) // file, os.Stdin
	scanner.Buffer(buf, 500100)

	var n int // 4<=n<=500'000
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	scanner.Scan()
	s := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	limit := n / 4
	ind := map[byte]int{'G': 0, 'A': 1, 'C': 2, 'T': 3}

	count := make([]GeneCount, n)

	upper := n
	l := 0
	count[l][ind[s[l]]]++
	for l = 1; l < n; l++ {
		count[l] = count[l-1]
		count[l][ind[s[l]]]++

		if upper == n && count[l][ind[s[l]]] > limit {
			upper = l
		}
	}

	total := count[n-1]
	min := n
	j_max := n

	// minimum length to be replaced
	repl := 0
	for p := 0; p < 4; p++ {
		if total[p] > limit {
			repl += total[p] - limit
		}
	}

	fmt.Println(upper, repl, upper+repl-1)

	c := 0
	// 0 <- i
	for i := upper - 1; i >= limit; i-- {

		j_max = i + min
		if j_max > n {
			j_max = n
		}

		// j -> n
		for j := i + repl; j < j_max; j++ {
			c++
			// check
			p := 0
			for p < 4 && count[i][p]+total[p]-count[j][p] <= limit {
				p++
			}

			if p >= 4 && min > j-i {
				min = j - i
				//fmt.Println(i, j, min)
				break
			}
		}
	}

	fmt.Println(min)
}
