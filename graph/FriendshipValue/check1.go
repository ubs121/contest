package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Student struct {
	x     int   // student id
	y     []int // friends (always x > y[i])
	group int
}

type Data []Student
type ByLen []Student

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i].y) < len(a[j].y) }

func main() {
	file, _ := os.Open("input1.txt")
	defer file.Close()

	reader := bufio.NewReader(file) // file, os.Stdin

	var q, n, m int

	fmt.Fscanf(reader, "%d\n", &q)

	for T := 0; T < q; T++ {
		fmt.Fscanf(reader, "%d %d\n", &n, &m)

		// do input
		var x, y int
		data := make(Data, n+1)

		for i := 0; i < m; i++ {
			fmt.Fscanf(reader, "%d %d\n", &x, &y)

			if x < y {
				x, y = y, x
			}

			// append 'y' into node 'x', always x > y
			data[x].y = append(data[x].y, y)
			data[x].x = x

		}

		sort.Sort(ByLen(data))

		st_group := make(map[int]int, n+1) // group registry (group[student] = group id)
		g_id := 0                          // group id
		var friends []int                  // friends

		// хамгийн цөөн найзтайгаас нь эхлэх
		for i := 0; i < len(data) && data[i].x != 0; i-- {
			x = data[i].x
			friends = data[i].y

			// өмнөх группуудын аль нэгтэй холбогдох уу?
			for j := 0; j < len(friends); j++ {
				y = friends[j]

				// (x, y) ?  сагсанд хийх үү?

			}
		}

		fmt.Println(data)
	}

}
