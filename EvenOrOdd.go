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