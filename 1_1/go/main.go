package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

func countIntInstances(intArray []int, numberToCount int, resultChan chan <- int,  wg *sync.WaitGroup) {
    defer wg.Done()
    multiplier := 0
    for _, val := range intArray {
        if val == numberToCount {
            multiplier += 1
        }
    }

    resultChan <- numberToCount * multiplier
    return
}

func consumePartialSums(resultChan chan int) int {
    total := 0
    for partialSum := range resultChan {
       total += partialSum 
    }
    
    return total
}

func calculateSimilarityScore(col_a []int, col_b []int) int {
    resultChan := make(chan int)
    var wg sync.WaitGroup

    for _, elem := range col_a {
        wg.Add(1)
        go countIntInstances(col_b, elem, resultChan, &wg) 
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    total := consumePartialSums(resultChan)
    return total
}

func main() {
    col_a, col_b, err := readColumnsFromFile() 
    if err != nil {
        log.Fatal("Oops!")
    }
    
    fmt.Println(calculateSimilarityScore(col_a, col_b))
}
