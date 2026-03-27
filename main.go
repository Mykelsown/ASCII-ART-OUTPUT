package main

import (
	asciiart "asciiartoutput/MethodsAndTesting"
	"flag"
	"fmt"
	"os"
)

func main() {
	var fileName string
	flagType := os.Args[1]

	flag.StringVar(&fileName, "output", "result", "The file where the ascii art is printed to")

	flag.Parse()
	arguments := flag.Args()
	formatType := arguments[len(arguments)-1]

	// This checks for wrong flag input the user might pass in through the terminal
	if flagType != "--output="+fileName {
		fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard`)
		return
	}

	contentRead, readinStatus := asciiart.FileHandler(fileName, formatType)
	for _, str := range arguments[:len(arguments)-1] {
		fmt.Println(asciiart.FormatPrinter(str, string(contentRead), readinStatus))
	}

}
