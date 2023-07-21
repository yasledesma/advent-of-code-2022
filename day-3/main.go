package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    var priorityOne int
    var priorityTwo int
    var elfs = make([]string, 0, 3)    
    
    input := OpenFile("input.txt")
    
    for input.Scan() {
        line := input.Text()
        
        priorityOne += CalcPartOne(line)

        elfs = append(elfs, line)

        if len(elfs) == 3 {
            priorityTwo += CalcPartTwo(elfs[0], elfs[1], elfs[2])
            elfs = elfs[:0]
        }
    }

    fmt.Print(priorityOne, priorityTwo)
}

func OpenFile(input string) *bufio.Scanner {
    file, err := os.Open(input)
    
    if err != nil {
        log.Fatal(err)
    }

    return bufio.NewScanner(file)
}

func CalcPartOne(line string) int {
    var points int
    letters := "abcdefghijklmnopqrstuvwxyz"
    
    middleIndex := len(line) / 2
    
    for _, v := range line[:middleIndex] {
        if strings.Contains(line[middleIndex:], string(v)) {
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

func CalcPartTwo(lineOne, lineTwo, lineThree string) int {
    var points int
    letters := "abcdefghijklmnopqrstuvwxyz"
    
    for _, v := range lineOne {
        if strings.Contains(lineTwo, string(v)) {
            if strings.Contains(lineThree, string(v)){
                if strings.Index(letters, string(v)) != -1 {
                    points += strings.Index(letters, string(v)) + 1
                } else {
                    points += strings.Index(letters, strings.ToLower(string(v))) + 27 
                }
                break
            }
        }
    }
    
    return points 
}
