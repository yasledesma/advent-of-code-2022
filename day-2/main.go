package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Link: https://adventofcode.com/2022/day/2

// Options:
// Column 1: A (Rock), B (Paper), C (Scissors)
// Column 2: X (Rock), Y (Paper), Z (Scissors)

// Points:
// Shape: Rock = 1, Paper = 2, Scissors = 3.
// Outcome: Win = 6, Draw = 3, Loss = 0.

// Round of three. What would your total score be if everything goes exactly according to your strategy guide?

func main() {
    var score int
    
    letters := map[string]map[string]int{
        "X": {"A": 0, "B": -1, "C": 1},
        "Y": {"A": 1, "B": 0, "C": -1},
        "Z": {"A": -1, "B": 1, "C": 0},
    }

    tools := map[string]int{
        "X": 1,
        "Y": 2,
        "Z": 3,
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

        switch letters[me][oponent] {
        case 1:
            score += 6 + tools[me]
        case -1:
            score += 0 + tools[me]
        default:
            score += 3 + tools[me]
        }
    }

    fmt.Println(score)
}

