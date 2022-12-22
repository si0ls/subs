package main

import (
	"fmt"
	"os"

	"github.com/si0ls/subs/stl"
)

func main() {
	// Get filepath from args
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filepath>")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse file
	s := stl.NewFile()
	if warns, err := s.Decode(file); err != nil {
		panic(err)
	} else if len(warns) > 0 {
		fmt.Println("Warnings:")
		printErrs(warns...)
		fmt.Println("====================================")
	}

	// Print file
	stl.PrintGSI(s.GSI)
	for _, tti := range s.TTI {
		stl.PrintTTI(tti, s.GSI.CCT)
	}
	fmt.Println("====================================")

	// Validate file
	if warns, err := s.Validate(); err != nil {
		panic(err)
	} else if len(warns) > 0 {
		fmt.Println("Warnings:")
		printErrs(warns...)
		fmt.Println("====================================")
	}

	// Encode XML to file
	xmlFile, err := os.Create("out.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	if err := s.EncodeXML(xmlFile); err != nil {
		panic(err)
	}

}

func printErrs(errs ...error) {
	for _, err := range errs {
		if err == nil {
			continue
		}
		fmt.Printf("[Err]: %s\n", err)
	}
}
