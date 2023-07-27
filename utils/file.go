package utils

import (
	"bufio"
	"os"
    "log"
)

func OpenFile(file string) *bufio.Scanner {
    input, err := os.Open(file)

    if err != nil {
        log.Fatal(err)
    }
    
    return bufio.NewScanner(input)
}
