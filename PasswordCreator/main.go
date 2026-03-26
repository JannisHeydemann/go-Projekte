package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type password struct {
	length          int
	include_numbers bool
	include_special bool
}

// Usage <PasswordGen> length numbers:true:optional specialsymbols:true:optional
func main() {

	alphabet_lower := "abcdefghijklmnopqrstuvwxyz"
	alphabet_upper := strings.ToUpper(alphabet_lower)
	special_chars := "!@#$%&*()_-+=?/><[]{}~"
	numbers := "123456789"

	// get Args
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Fprintln(os.Stderr, "Usage: <myprgram> Length:Number IncludeNumbers:Bool IncludeSpecialSymbols:Bool")
		os.Exit(0)
	}

	password_length := os.Args[1]

	var include_numbers string
	var include_special string

	if len(os.Args) > 2 {
		include_numbers = os.Args[2]
		include_special = os.Args[3]
	} else {
		include_numbers = "true"
		include_special = "true"
	}

	// Create a string of all usable characters
	usable := alphabet_lower + alphabet_upper

	if include_numbers != "false" {
		usable += numbers
	}
	if include_special != "false" {
		usable += special_chars
	}

	// Convert the usable string to an array
	var sub_usable []string
	for _, r := range usable {
		sub_usable = append(sub_usable, string(r))
	}

	password_length_int, err := strconv.Atoi(password_length)
	if err != nil {
		panic(err)
	}

	var password string

	// Loop through the password length
	for i := 0; i < password_length_int; i++ {
		password += sub_usable[rand.Intn(len(sub_usable))]
	}

	// Output the password
	fmt.Println("The finished password is: " + password)

}
