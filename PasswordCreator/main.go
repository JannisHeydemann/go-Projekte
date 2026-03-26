package PasswordCreator

import (
	"fmt"
	"os"
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

}
