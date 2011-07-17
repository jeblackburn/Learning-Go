package main

import (
"fmt"
"os"
"strings"
)

func main() {
    greeting := "world!"
    fmt.Println("len(os.Args) = ", os.Args[0])
    if len(os.Args) > 1 {
        greeting = strings.Join(os.Args[1:], " ")
    }
    fmt.Println("Hello,", greeting)
}
