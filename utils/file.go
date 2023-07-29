package utils

import (
	"bufio"
	"os"
    "log"
)

func OpenFile(file string) (*bufio.Scanner, *os.File) {
    input, err := os.Open(file)

    LogError(err)
    
    return bufio.NewScanner(input), input
}

func LogError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
