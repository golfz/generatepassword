package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	defaultLength = 16
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

	rand.Seed(time.Now().UnixNano())
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

	fmt.Println(generatePassword())
}

func generatePassword() string {
	var password []byte

	source := ""

	if !noUpper {
		source += uppers
	}
	if !noLower {
		source += lowers
	}
	if !noNumber {
		source += numbers
	}
	if !noSpecial {
		source += specials
	}

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(source))
		password = append(password, source[randNum])
	}

	return string(password)
}
