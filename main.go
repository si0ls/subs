package main

import (
	"fmt"
	"os"

	"github.com/si0ls/go-subs/stl"
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

	// Parse file
	s := stl.NewFile()
	if warns, err := s.Decode(file); err != nil {
		panic(err)
	} else if len(warns) > 0 {
		fmt.Println("Warnings:")
		for _, warn := range warns {
			fmt.Println(warn)
		}
		fmt.Println("====================================")
	}

	// Print file
	stl.PrintGSI(s.GSI)
	for _, tti := range s.TTI {
		stl.PrintTTI(tti, s.GSI.CCT)
	}
}
