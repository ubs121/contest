package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	v   int   // student id
	adj []int // friends

	// BFS data
	parent int // parent
	color  int // White=0, Gray=1, Black=2
	d      int // distance from source/root
}
type Data []Student

// do BFS
func total_BFS(data Data, root int) int64 {
	n := 1
	repeated := 0

	var Q []int
	var fs []int
	var u, v int

	Q = append(Q, root)
	for len(Q) > 0 {
		// dequeue
		u = Q[0]
		Q = Q[1:]

		fs = data[u].adj // friends

		for i := 0; i < len(fs); i++ {
			v = fs[i]

			if data[v].color == 0 {
				data[v].color = 1
				//data[v].d = data[u].d + 1
				//data[v].parent = u

				// enqueue
				Q = append(Q, v)

				n += 1
			} else if data[v].color == 1 {
				// repeated link
				repeated += 1

				//fmt.Println(u, fs[v])
			}
		}

		data[u].color = 2
	}

	// calculate the total
	total := int64((n-1)*(n)*(n+1)) / 3
	total = total + int64(n*(n-1)*repeated)

	return total
}

func main() {
	file, _ := os.Open("input00.txt")
	defer file.Close()

	reader := bufio.NewReader(file) // file, os.Stdin

	var q int
	fmt.Fscanf(reader, "%d\n", &q)

	// iterate over test cases
	for T := 0; T < q; T++ {
		var n, m int
		fmt.Fscanf(reader, "%d %d\n", &n, &m)

		// do input
		var x, y int
		data := make(Data, n+1)

		for i := 0; i < m; i++ {
			fmt.Fscanf(reader, "%d %d\n", &x, &y)

			//data[x].v = x
			data[x].adj = append(data[x].adj, y)

			//data[y].v = y
			data[y].adj = append(data[y].adj, x)
		}

		var total, total_max int64

		// choose root, FIXME: optimize ?!
		for root := 1; root < n+1; root++ {
			if data[root].color == 0 {
				total = total_BFS(data, root)

				if total > total_max {
					total_max = total
				}
			}
		}

		fmt.Println(total_max)

	}

}
