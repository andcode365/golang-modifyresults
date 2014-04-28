package main
import "fmt"

func d1() (result int) {
    defer func() {
        // do nothing
    }()
    return 0 
}

func d2() (result int) {
    defer func() {
        result++
    }()
    return 0 
}

func d3() (result int) {
    defer func() {
        result++
    }()
    result = 1
    return 0 
}

func d4() (result int) {
    defer func() int {
        result++
        return 0
    }()
    return 0 
}

func main() {
    fmt.Println("d1=", d1())
    fmt.Println("d2=", d2())
    fmt.Println("d3=", d3())
    fmt.Println("d4=", d4())
}
