package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
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

func CLOC(filename string) (int, int) {
    lines, err := ReadLines(filename)
    if err != nil {
        fmt.Println("[ERROR]: bufio.NewScanner->readLines() failed.")
        os.Exit(0)
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
