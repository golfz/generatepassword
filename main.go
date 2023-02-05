package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultLength = 8
	uppers        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowers        = "abcdefghijklmnopqrstuvwxyz"
	numbers       = "0123456789"
	specials      = "!@#$%^&*()_+"
)

var (
	length    int
	noUpper   bool
	noLower   bool
	noNumber  bool
	noSpecial bool
)

func init() {
	length = defaultLength
}

func main() {
	// create -h flag to print help
	for i, arg := range os.Args {
		if arg == "-h" {
			fmt.Println("Usage: go run main.go [-l <length>]")
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
