package main

import (
	"fmt"
	"os"
    "octal/utils"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: octal <file>")
        os.Exit(0)
    }
    loc, empty := utils.CLOC(os.Args[1])
    fmt.Println(os.Args[1:])
    fmt.Printf("LOC: %d\n", loc)
    fmt.Printf("Empty: %d\n", empty)
    fmt.Printf("Total: %d\n", loc + empty)
}
