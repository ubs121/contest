package main

import "fmt"
import "math"

func main() {
    var n int
    fmt.Scanf("%d", &n)

    // read array
    arr := [][]int{}
    arr = make([][]int, n)

    for i := 0; i < n; i++ {
        arr[i] = make([]int, n)

        for j := 0; j < n; j++ {
            fmt.Scanf("%d", &arr[i][j])
        }
    }

    // print array
    s1 := 0
    s2 := 0
    for i := 0; i < n; i++ {
        s1 += arr[i][i]     // diag 1
        s2 += arr[i][n-i-1] // diag 2
    }

    fmt.Println(math.Abs(float64(s2 - s1)))

}
