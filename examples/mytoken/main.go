package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Bios-Marcel/discordgo"
)

// Variables used for command line parameters
var (
	Email    string
	Password string
	SecondF  string
)

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&SecondF, "t", "", "MFA OTP")
	flag.Parse()

	if Email == "" || Password == "" || SecondF == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	// Create a new Discord session using the provided login information.
	dg, err := discordgo.NewWithPasswordAndMFA(Email, Password, SecondF)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Print out your token.
	fmt.Printf("Your Authentication Token is:\n\n%s\n", dg.Token)
}
