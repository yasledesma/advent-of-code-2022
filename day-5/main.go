package main

import (
	"fmt"
	"strings"

	"github.com/yasledesma/go-advent-of-code-2022/utils"
)

// Load the first 9 lines of input into memory, then find out at what line do the instructions start and begin loading them and processing them one by one.
// 8x9 matrix

func main() {
    var move [3]int
    stack := map[int]string{}
    
    file := utils.OpenFile("input.txt")

    for file.Scan() {
        i := 4
        line := file.Text()

        if strings.Contains(line, "[") {
            for j := 0; j < len(line); j++ {
                if string(line[j]) != " " && string(line[j]) != "[" && string(line[j]) != "]" {
                    stack[(i-1)/4] += string(line[j])
                }
                i++
            }
        }

        if strings.Contains(line, "move") {
            fmt.Sscanf(line, "move %d from %d to %d", &move[0], &move[1], &move[2])
            
            for j := 0; j < move[0]; j++ {
                rest := string(stack[move[1]][1:])
                letter := string(stack[move[1]][0])
                
                stack[move[1]] = rest
                stack[move[2]] = letter + stack[move[2]]
            }
        }
    }
    
    fmt.Println(stack)
}


