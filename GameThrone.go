// https://www.hackerrank.com/challenges/game-of-thrones
package main

import "fmt"

func isPalindrom(s string) bool {
    var letters [26]int

    for i := 0; i < len(s); i++ {
        letters[s[i]-'a']++
    }

    //fmt.Println(letters)

    odd := 0
    for i := 0; i < len(letters); i++ {
        if letters[i]%2 == 1 {
            odd++

            if odd > 1 {
                return false
            }
        }
    }

    return true
}

func main() {
    var s string
    fmt.Scanln(&s)

    if isPalindrom(s) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }

}
