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
			help := `Usage: getpass [-l <length>] [-no-upper] [-no-lower] [-no-number] [-no-special]
    -l <int>      length of password (default 16)
    -no-upper     exclude upper case letters
    -no-lower     exclude lower case letters
    -no-number    exclude numbers
    -no-special   exclude special characters`
			fmt.Println(help)
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
		if arg == "-no-upper" {
			noUpper = true
		}
		if arg == "-no-lower" {
			noLower = true
		}
		if arg == "-no-number" {
			noNumber = true
		}
		if arg == "-no-special" {
			noSpecial = true
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
