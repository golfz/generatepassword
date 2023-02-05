package main

import (
	"fmt"
	"github.com/golfz/getpass/password"
	"os"
	"strconv"
)

var (
	length    int
	noUpper   bool
	noLower   bool
	noNumber  bool
	noSpecial bool
)

func init() {
	length = password.DefaultLength
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

	fmt.Println(password.GeneratePassword(length, !noUpper, !noLower, !noNumber, !noSpecial))
}
