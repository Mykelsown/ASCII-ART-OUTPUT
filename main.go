package main

import (
	asciiart "asciiartoutput/MethodsAndTesting"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	flagType := os.Args[1]

	flag.StringVar(&fileName, "output", "result", "The file where the ascii art is printed to")

	flag.Parse()
	arguments := flag.Args()
	formatType := arguments[len(arguments)-1]

	// This checks for wrong flag input the user might pass in through the terminal
	contentRead, readingStatus := asciiart.FileHandler(fileName, formatType)
	if flagType == "--output="+fileName && fileName != "white" {
		for _, str := range arguments[:len(arguments)-1] {
			writingErr := os.WriteFile(fileName, []byte(asciiart.FormatPrinter(str, string(contentRead), readingStatus)), 0666)
			if writingErr != nil {
				log.Fatalf("err!: %W", writingErr)
			}
		}
	} else if strings.HasPrefix(flagType, "-output=") || strings.HasPrefix(flagType, "output=") || strings.HasPrefix(flagType, ".output=") {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"`)
		return
	} else {
		fmt.Println(asciiart.FormatPrinter(arguments[len(arguments)-1], string(contentRead), readingStatus)) // this allows for validation of the base ascii-art project
	}

}
