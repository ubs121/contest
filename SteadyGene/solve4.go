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
	file, _ := os.Open("input07.txt")
	defer file.Close()

	buf := make([]byte, 500100)
	scanner := bufio.NewScanner(os.Stdin) // file, os.Stdin
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

	// calculate left part
	left := make([]GeneCount, n)
	left_max := n
	l := 0
	left[l][ind[s[l]]]++
	for l = 1; l < n; l++ {
		left[l] = left[l-1]
		left[l][ind[s[l]]]++

		if left_max == n && left[l][ind[s[l]]] > limit {
			left_max = l
		}
	}

	// [left]+[middle]+[right]=[total]
	total := left[n-1]

	over := 0             // minimum length to be replaced
	var mid_max GeneCount // each minimum letters to be replaced
	for p := 0; p < 4; p++ {
		mid_max[p] = total[p] - limit

		// over sized letters
		if total[p] > limit {
			over += mid_max[p]
		}

	}

	c := 0
	j_max := n
	min := n // should be min >= over
	var middle, right GeneCount
	var i, j, p int

	fmt.Printf("left_max=%v, total=%v, to_repl=%v\n", left_max, total, mid_max)

	// 0 <- i
	for i = left_max - 2; i >= 0; i-- {
		c++
		// calculate j_max
		j_max = i + min
		if j_max > n {
			j_max = n
		}

		// init j
		j = i + over

		// count skipped letters
		for p := 0; p < 4; p++ {
			middle[p] = left[j][p] - left[i][p]
			right[p] = total[p] - left[j][p]
		}

		// j -> n
		for j < j_max {

			c++

			p = ind[s[j]]

			if middle[p] >= mid_max[p] && mid_max[p] > 0 {
				fmt.Println("-----------")
				break
			}

			fmt.Println(left[i], middle, right)

			middle[p]++
			right[p]--
			j++
		}

		if min > j-i {
			min = j - i
		}

	}

	fmt.Println(min, c)
}
