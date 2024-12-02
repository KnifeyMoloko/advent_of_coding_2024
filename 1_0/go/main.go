package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "errors"
    "strconv"
    "sort"
    "math"
)

func readColumnsFromFile() (col_a []int, col_b []int, err error) {
    col_a = []int{}
    col_b = []int{}
    content, err := os.ReadFile("input")
    if err != nil {
        log.Print("Failed to read file")
        return col_a, col_b, errors.New("Could not read from file.")
    }

    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        split_line := strings.Split(line, "  ")
        if len(split_line) == 2 {
            int_a, err := strconv.Atoi(strings.TrimSpace(split_line[0]))
            if err != nil {
                log.Fatalf("Failed to convert first string to int. String:%s", split_line[0])
            }
            
            int_b, err := strconv.Atoi(strings.TrimSpace(split_line[1]))
            if err != nil {
                log.Fatalf("Failed to convert second string to int. String:%s", split_line[1])
            }
            col_a = append(col_a, int_a)
            col_b = append(col_b, int_b)
        }
    }
    return col_a, col_b, nil
}

func sumDiffs(col_a []int, col_b []int) int {
    diffSum := 0.0
    sort.Slice(
        col_a, func(i, j int) bool {
            return col_a[i] < col_a[j]
        },
    )
    sort.Slice(
        col_b, func(i, j int) bool {
            return col_b[i] < col_b[j]
        },
    )

    for i, val := range col_a {
       diffSum += math.Abs(float64(val) - float64(col_b[i])) 
    }

    return int(diffSum)
}

func main() {
    col_a, col_b, err := readColumnsFromFile() 
    if err != nil {
        log.Fatal("Oops!")
    }
    diffSum := sumDiffs(col_a, col_b)
    fmt.Printf("Output: %d\nBye bye!\n", diffSum)
}
