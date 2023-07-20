package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Link: https://adventofcode.com/2022/day/2
// Round of three. What would your total score be if everything goes exactly according to your strategy guide?

func main() {
    var score int
    
    outcome := map[string]map[string]string{
        "X": {"A": "S", "B": "R", "C": "P"}, // Loss: Oponent wins against...
        "Y": {"A": "R", "B": "P", "C": "S"}, // Draw: Oponent draws against...
        "Z": {"A": "P", "B": "S", "C": "R"}, // Win: Oponent losses against...
    }

    
    result := map[string]int{
        "X": 0, // Loss = 0
        "Y": 3, // Draw = 3
        "Z": 6, // Win = 6
    }
    
    options := map[string]int{
        "R": 1, // Rock = 1
        "P": 2, // Paper = 2
        "S": 3, // Scissors = 3
    }

    file, err := os.Open("input.txt")
    
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    input := bufio.NewScanner(file)

    for input.Scan() {
        oponent := string(input.Text()[0])
        me := string(input.Text()[2])
        
        score += result[me] + options[outcome[me][oponent]]
    }

    fmt.Println(score)
}

