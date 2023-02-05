package main

import (
	"os"
	"strconv"
)

const (
	defaultLength = 8
)

var (
	length int
)

func init() {
	length = defaultLength
}

func main() {
	// create -h flag to print help
	for i, arg := range os.Args {
		if arg == "-h" {

			return
		}
		if arg == "-l" {
			sL := os.Args[i+1]
			iL, err := strconv.Atoi(sL)
			if err != nil {
				panic("Invalid length")
			}
			length = iL
		}
	}
}
