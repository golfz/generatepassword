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
			printHelp()
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

func printHelp() {
	fmt.Println("Usage: getpass [-l <length>] [-no-upper] [-no-lower] [-no-number] [-no-special] [-h]")
	fmt.Println("    -l <int>      length of password (default 16)")
	fmt.Println("    -no-upper     exclude upper case letters")
	fmt.Println("    -no-lower     exclude lower case letters")
	fmt.Println("    -no-number    exclude numbers")
	fmt.Println("    -no-special   exclude special characters")
	fmt.Println("    -h            show help")
}
