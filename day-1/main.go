package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Link: https://adventofcode.com/2022/day/1
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

// 1. Read the input file and get the data out from it.
// 2. Split the input in lines by it's blank spaces.
// 3. Sum the numbers.
// 4. Get the largest one.

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var current int64
    var biggest [3]int64

    for scanner.Scan() {
        if scanner.Text() != "" {
            num, err := strconv.ParseInt(scanner.Text(), 10, 64)
            
            if err != nil {
                log.Fatal(err)
            }

            current = current + num
        } else {
            if current > biggest[0] {
                biggest[0] = current
            } else if current > biggest[1] {
                biggest[1] = current
            } else if current > biggest[2] {
                biggest[2] = current
            }

            current = 0
        }
    }

    fmt.Println(biggest[0] + biggest[1] + biggest[2])
}
