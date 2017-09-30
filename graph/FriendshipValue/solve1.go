package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	x, y int
}

type MemberCount struct {
	n        int // +1 link
	repeated int // repeated pair
}

type ByX []Pair

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].x < a[j].x }

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
		pairs := make([]Pair, m)
		children := make([]int, n+1)

		for i := 0; i < m; i++ {
			fmt.Fscanf(reader, "%d %d\n", &x, &y)

			if x < y {
				pairs[i].x = x
				pairs[i].y = y
			} else {
				pairs[i].x = y
				pairs[i].y = x
			}

			children[x]++
		}

		sort.Sort(ByX(pairs))

		st_group := make([]int, n+1)             // student group registry (group[student] = group id)
		member := make(map[int]MemberCount, n+1) // member count for a group ( group_id -> (n, repeated) )
		g_id := 0                                // group id

		var fc MemberCount // friend count

		for p := 0; p < m; p++ {
			x = pairs[p].x
			y = pairs[p].y

			if st_group[x] > 0 {
				fc = member[st_group[x]]

				if st_group[y] > 0 {
					if st_group[x] != st_group[y] {
						fc.n += member[st_group[y]].n
						fc.repeated += member[st_group[y]].repeated

						// merge groups
						for i := 1; i <= g_id; i++ {
							if st_group[y] == st_group[i] {
								st_group[i] = st_group[x]
							}
						}

						// delete st_group[y]
						delete(member, st_group[y])
					} else {
						// same group, push into stack, add at group end
						fc.repeated += 1

					}
				} else {
					st_group[y] = st_group[x] // put y into x's group
					fc.n += 1                 // +1 new member
				}

				// update member count
				member[st_group[x]] = fc
			} else {
				if st_group[y] > 0 {
					st_group[x] = st_group[y] // put x into y's group

					fc = member[st_group[y]]
					fc.n += 1 // +1 new member
					member[st_group[y]] = fc
				} else {
					g_id = g_id + 1                  // create a new group
					member[g_id] = MemberCount{2, 0} // 2 new member

					st_group[x] = g_id
					st_group[y] = g_id
				}
			}

		}

		fmt.Println(children)
		//fmt.Println(pairs)

		// calculate totals
		var total, max int64
		for _, fc := range member {
			total = int64((fc.n-1)*(fc.n)*(fc.n+1)) / 3
			total = total + int64(fc.n*(fc.n-1)*fc.repeated)

			if max < total {
				max = total
			}

			//fmt.Println(g_id, "=", total)
		}

		fmt.Println(max)
	}

}
