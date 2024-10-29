package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func count(filename string) (int, int) {
    lines, err := readLines(filename)
    if err != nil {
        fmt.Println("[ERROR]: bufio.NewScanner->readLines() failed.")
    }

    var loc int = 0
    var empty int = 0

    for i := 0; i < len(lines); i++ {
        var line string = lines[i]
        if len(strings.TrimSpace(line)) == 0 {
            empty++
            continue
        }
        loc++
    }

    return loc, empty
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("usage: octal <file>")
        os.Exit(0)
    }
    loc, empty := count(os.Args[1])
    fmt.Println(os.Args[1:])
    fmt.Printf("LOC: %d\n", loc)
    fmt.Printf("Empty: %d\n", empty)
    fmt.Printf("Total: %d\n", loc + empty)
}
