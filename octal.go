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
    } else if os.Args[1] == "--help" || os.Args[1] == "-h" {
        fmt.Printf("usage: octal [options] <file>\n")
        fmt.Printf("-h, --help: Brings up this message\n")
        os.Exit(0)
    }
    loc, empty := utils.CLOC(os.Args[1])
    fmt.Println(os.Args[1:])
    fmt.Printf("LOC: %d\n", loc)
    fmt.Printf("Empty: %d\n", empty)
    fmt.Printf("Total: %d\n", loc + empty)
}
