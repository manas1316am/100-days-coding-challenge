package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwr := []rune(pw)
	if len(pwr) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwr {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("password is good")
	} else {
		fmt.Println("password is bad")
	}
	if passwordChecker("This!I5A") {
		fmt.Println("password is good")
	} else {
		fmt.Println("Password is bad")
	}
}
