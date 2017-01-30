package main

import "fmt"
import "os"
import "bufio"

var pow2 map[int]int

// calculate 2 powers
func init_pow() {
	pow2 = make(map[int]int)

	p := 1
	for i := 0; i < 65; i++ {
		pow2[i] = p
		p = p << 1
	}
}

func xorCount(x int) int {
	i := x

	// number of binary digits
	d := 0
	for i > 0 {
		d += 1
		i = i / 2
	}

	return pow2[d] - x - 1
}

func main() {

	file, _ := os.Open("input08.txt")
	defer file.Close()

	reader := bufio.NewReader(file) // file, os.Stdin

	var q, x int
	fmt.Fscanf(reader, "%d\n", &q)

	init_pow()
	for i := 0; i < q; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		fmt.Println(xorCount(x))
	}
}
