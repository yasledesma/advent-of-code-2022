package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    var score int
    input := OpenFile("input.txt")
    
    for input.Scan() {
        line := input.Text()
        middleIndex := len(line) / 2
        
        score += CalculateItems(line[:middleIndex], line[middleIndex:])
    }

    fmt.Print(score)
}

func OpenFile(input string) *bufio.Scanner {
    file, err := os.Open(input)
    
    if err != nil {
        log.Fatal(err)
    }

    return bufio.NewScanner(file)
}

func CalculateItems(comp1, comp2 string) int {
    var points int
    letters := "abcdefghijklmnopqrstuvwxyz"
    
    for _, v := range comp1 {
        if strings.Contains(comp2, string(v)) {
            if strings.Index(letters, string(v)) != -1 {
                points += strings.Index(letters, string(v)) + 1
            } else {
                points += strings.Index(letters, strings.ToLower(string(v))) + 27 
            }
            break
        }
    }
    
    return points 
}
