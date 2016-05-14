package main

import "fmt"

func jump(river string, pos int, vel int) bool {
	if pos >= len(river) {
		return true
	}

	if river[pos] == 'W' {
		return false
	}

	if jump(river, pos+vel, vel) {
		return true
	}

	if jump(river, pos+vel+1, vel) {
		return true
	}

	if jump(river, pos+vel-1, vel) {
		return true
	}

	return false
}

func main() {
	river := "RWWRW"
	pos := -1
	vel := 2

	if jump(river, pos+vel, vel) {
		fmt.Println("YES")
	} else if jump(river, pos+vel+1, vel) {
		fmt.Println("YES")
	} else if jump(river, pos+vel-1, vel) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
