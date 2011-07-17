package main

import (
"fmt"
"os"
"strings"
)

func main() {
    greeting := "world!"
    if len(os.Args) > 1 {
        greeting = strings.Join(os.Args[1:], " ")
    }
    fmt.Println("Hello,", greeting)
}
