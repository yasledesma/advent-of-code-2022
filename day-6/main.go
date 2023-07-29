package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    var markOne = 4
    var startOne = 0
    
    file, err := os.ReadFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    // PART 1 
    for {
        current := file[startOne:markOne]
        
        if current[3] != current[2] && current[3] != current[1] && current[3] != current[0] && current[2] != current[1] && current[2] != current[0] && current[1] != current[0] {
            break
        }

        markOne++
        startOne++
    }

    // PART 2
    var markTwo = 14 
    var startTwo = 0
    var stream = make(map[byte]string)
    var currenTwo = file[startTwo:markTwo]
    j := 0
    
    for {
        if len(stream) == 14 {
            break
        } else if stream[currenTwo[j]] == "" {
            stream[currenTwo[j]] = fmt.Sprint(j)
            j++
        } else {
            stream = map[byte]string{}
            startTwo++
            markTwo++
            currenTwo = file[startTwo:markTwo]
            j = 0
        }
    }

    
    fmt.Println(markOne, markTwo)
}
