package main

import (
	"fmt"
	"google-two-factor-auth/googleauth"
	"os"
	"time"
)

// all []byte in this program are treated as Big Endian
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must specify key to use")
		os.Exit(1)
	}

	input := os.Args[1]
	epochSeconds := time.Now().Unix()
	
	pwd, err := googleauth.GetGoogleAuth(input, epochSeconds)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	secondsRemaining := 30 - (epochSeconds % 30)
	var userAccount  = "x893675@example.org"
	var companyName  = "example"
	var algo  = "SHA1"
	var digits  = "6"
	var period  = "30"
	fmt.Println(googleauth.GenGoogleAuthTotpUri(input, userAccount, companyName, algo, digits, period))
	fmt.Printf("%s (%d second(s) remaining)\n", pwd, secondsRemaining)
}
