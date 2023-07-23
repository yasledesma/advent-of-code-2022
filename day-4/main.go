package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    var fullOverlaps int
    var partialOverlaps int
    
    input := OpenFile("input.txt")

    for input.Scan() {
        line := input.Text()
        commaIndex := strings.Index(line, ",")
        hyphenIndexFirst := strings.Index(line[:commaIndex], "-")
        hyphenIndexLast := strings.Index(line[commaIndex + 1:], "-")
        slotOneOne := line[:commaIndex][:hyphenIndexFirst]
        slotOneTwo := line[:commaIndex][hyphenIndexFirst + 1 :]
        slotTwoOne := line[commaIndex + 1:][:hyphenIndexLast]
        slotTwoTwo := line[commaIndex + 1:][hyphenIndexLast + 1:]
        
        // How is this the right answer??? PART 1
        if (slotTwoOne >= slotOneOne && slotOneOne <= slotTwoTwo) || (slotTwoOne >= slotOneTwo && slotOneTwo <= slotTwoTwo) {
            fullOverlaps += 1
        }

        // PART 2: How tf do I keep getting "the right answer for someone else" this day??? This code doesn't make any sense either!
        // Am I crazy or dumb?
        if (slotOneOne <= slotTwoOne || slotOneOne <= slotTwoTwo) {
            partialOverlaps += 1
        }
    }

    fmt.Print(fullOverlaps, partialOverlaps)
}

func OpenFile(file string) *bufio.Scanner {
    input, err := os.Open(file)

    if err != nil {
        log.Fatal(err)
    }
    
    return bufio.NewScanner(input)
}
