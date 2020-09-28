package main

import (
	"fmt"
	"strconv"
	// color "github.com/bclicn/color"
)

// not gonna render a menu, it'll all be

// Menu Renders a menu that the user can interact with
func Menu() {
	// need to add param for the count of available JDKs to choose from
	var input string

	var jdk int
	// var lib int

	fmt.Print(au.Bold(au.Cyan(" ::")))
	fmt.Println(au.Bold(" There are 15 available JDKs to choose from."))
	fmt.Print(au.Bold(au.Cyan(" ::")))
	fmt.Println(au.Bold(" Java JDK Implementations:"))
	fmt.Println("    1) Oracle OpenJDK   2) Amazon Corretto   3) BellSoft Liberica   4) AdoptOpenJDK")
	fmt.Print(au.Bold(au.Cyan(" ::")))
	fmt.Print(" Enter a selection (default=1): ")

	fmt.Scanln(&input)
	fmt.Println("you entered:", input)

	state, e := strconv.Atoi(input) // need to deal with error handling, but for now, ignore the error
	if e != nil {
		// do nothing for now
		// will call a function that can be recursively called, one that expects correct input
		// (unless i can figure out how to make this one recursive)
	}

	// fmt.Printf(color.BPurple(" ::"))

	switch state {
	case openJDK:
		jdk = openJDK

	case corretto:
		jdk = corretto

	case liberica:
		jdk = liberica

	case adoptOpenJDK:
		jdk = adoptOpenJDK

	default:
		jdk = openJDK
	}
	fmt.Println("using", jdk)
}

// Help Print some helpful information
func Help() {
	// TODO: add a help printing method for program when a "-h" or "--help" arg is passed
	fmt.Printf("help to be implemented\n")
}

// Version Print Version information
func Version() {
	fmt.Println("version info to be made available soon-ish")
}
