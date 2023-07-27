package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    var mark = 4
    var start = 0
    
    file, err := os.ReadFile("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    // PART 1 
    for {
        current := file[start:mark]
        if current[3] != current[2] && current[3] != current[1] && current[3] != current[0] && current[2] != current[1] && current[2] != current[0] && current[1] != current[0]  {
            fmt.Println(mark)
            break
        }

        mark++
        start++
    }
}
