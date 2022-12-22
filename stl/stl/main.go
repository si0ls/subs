package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/si0ls/subs/stl"
	"os"
)

func main() {
	var inputFile string
	var outputFile string

	// flags declaration using flag package
	flag.StringVar(&inputFile, "input", "", "Specify input file.")
	flag.StringVar(&outputFile, "output", "", "Specify output path.")

	flag.Parse() // after declaring flags we need to call it

	fmt.Println(inputFile, outputFile)

	// Open file
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse file
	s := stl.NewFile()
	warns, err := s.Decode(file)
	if err != nil {
		panic(err)
	} else if len(warns) > 0 {
		fmt.Println("Warnings:")
		printErrs(warns...)
		fmt.Println("====================================")
	}

	// Validate file
	if warns, err := s.Validate(); err != nil {
		panic(err)
	} else if len(warns) > 0 {
		fmt.Println("Warnings:")
		printErrs(warns...)
		fmt.Println("====================================")
	}

	// Encode XML to file
	xmlFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	xmlFile.Write([]byte(xml.Header))
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

// go run monBinaire sources_100/1.stl -o toto.stl.xml
// go run monBinaire sources_100/1.stl -o toto.ttml
