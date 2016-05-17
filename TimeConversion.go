package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	hour := (s[0]-'0')*10 + (s[1] - '0')

	if s[8:10] == "PM" {
		if hour < 12 {
			hour = hour + 12
		}

	} else {
		if hour >= 12 {
			hour = hour - 12
		}
	}

	fmt.Printf("%02d:%s\n", hour, s[3:8])

}
