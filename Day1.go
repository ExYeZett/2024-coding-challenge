// Coding Challenge 2024
// 1/366
// https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

package kata

import "fmt"

func EvenOrOdd(n int) string {
    if n%2 == 0 {
        return "Even"
    }
    return "Odd"
}

func main() {
    fmt.Println(EvenOrOdd(7))
    fmt.Println(EvenOrOdd(10))
}