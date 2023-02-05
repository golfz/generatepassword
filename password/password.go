package password

import (
	"math/rand"
	"time"
)

const (
	DefaultLength = 16
	uppers        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowers        = "abcdefghijklmnopqrstuvwxyz"
	numbers       = "0123456789"
	specials      = "!@#$%^&*()-_+[:]"
)

func GeneratePassword(length int, useUpper, useLower, useNumber, useSpecial bool) string {
	var password []byte

	rand.Seed(time.Now().UnixNano())

	source := ""

	if useUpper {
		source += uppers
	}
	if useLower {
		source += lowers
	}
	if useNumber {
		source += numbers
	}
	if useSpecial {
		source += specials
	}

	for i := 0; i < length; i++ {
		randNum := rand.Intn(len(source))
		password = append(password, source[randNum])
	}

	return string(password)
}
