// https://www.hackerrank.com/challenges/sorted-subsegments
package main

import "fmt"
import "os"
import "bufio"

/**
* 	A -  array of integers
* 	l - left boundary
* 	r - right boundary
 */
func countingSort(A []int, l, r int) {
	n := r - l + 1
	B := make([]int, n)

	C := make(map[int]int)

	// [l,r] интервалд i хэдэн удаа байгааг тоолох
	for i := 0; i < n; i++ {
		C[A[l+i]] += 1
	}
	fmt.Println(C)

	// i-1 тоо хэдэн удаа байгааг тоолох
	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1] // FIXME: энд алдаа байна !!!
	}

	fmt.Println("DONE", len(C))

	// тоонуудыг байрлуулах
	for i := n - 1; i >= 0; i-- {
		B[C[A[l+i]]-1] = A[l+i]
		C[A[l+i]] = C[A[l+i]] - 1
	}

	// B -> A
	for i := 0; i < n; i++ {
		A[l+i] = B[i]
	}

	// DEBUG: fmt.Printf("(%d, %d): %v\n", l, r, A)
}

func main() {
	file, _ := os.Open("input11.txt")
	defer file.Close()

	reader := bufio.NewReader(file) // file, os.Stdin

	var n, q, k int

	fmt.Fscanf(reader, "%d %d %d\n", &n, &q, &k)

	A := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
	}

	// TODO:  оролтын өгөгдөлд сөрөг утга байна, эхлээд боловсруулалт хийх хэрэгтэй
	// дундаж утгыг олоод эерэг дараалал руу шилжүүлэх үү?

	var l, r int

	// q удаа эрэмбэлэх
	for i := 0; i < q; i++ {
		fmt.Fscanf(reader, "%d %d\n", &l, &r)
		fmt.Println("Sorting ", l, r)

		if l < r {
			countingSort(A, l, r)
		}

	}

	// k байрлал дахь тоог хэвлэх
	fmt.Printf("%d\n", A[k])

}
