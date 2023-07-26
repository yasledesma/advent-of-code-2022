// This is the shitties, most umperformant solution ever. I should have used channels in some way.
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
    stackOne := map[int]string{}
    stackTwo := map[int]string{}
    
    file := utils.OpenFile("input.txt")

    for file.Scan() {
        i := 4
        line := file.Text()

        if strings.Contains(line, "[") {
            for j := 0; j < len(line); j++ {
                if string(line[j]) != " " && string(line[j]) != "[" && string(line[j]) != "]" {
                    stackOne[(i-1)/4] += string(line[j])
                    stackTwo[(i-1)/4] += string(line[j])
                }
                i++
            }
        }

        if strings.Contains(line, "move") {
            fmt.Sscanf(line, "move %d from %d to %d", &move[0], &move[1], &move[2])
            
            // PART 1 PROCESS
            for j := 0; j < move[0]; j++ {
                restOne := string(stackOne[move[1]][1:])
                letter := string(stackOne[move[1]][0])
                
                stackOne[move[1]] = restOne
                stackOne[move[2]] = letter + stackOne[move[2]]
            }

            // PART 2 PROCESS
            restTwo := string(stackTwo[move[1]][move[0]:])
            letters := string(stackTwo[move[1]][0:move[0]])

            stackTwo[move[1]] = restTwo
            stackTwo[move[2]] = letters + stackTwo[move[2]]
        }
    }
    
    fmt.Println(stackOne)
    fmt.Println(stackTwo)
}


