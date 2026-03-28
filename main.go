package main

import (
	asciiart "asciiartoutput/MethodsAndTesting"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const usageMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard"

func main() {
	// Fix 5: guard against zero arguments before touching os.Args[1]
	if len(os.Args) < 2 {
		fmt.Println(usageMsg)
		return
	}

	var fileName string
	// Fix 3: default is "" so fileName != "" is an unambiguous "flag was set" check
	flag.StringVar(&fileName, "output", "", "The file where the ascii art is printed to")
	flag.Parse()

	arguments := flag.Args()

	// Fix 6: guard against empty arguments after flag parsing
	if len(arguments) == 0 {
		fmt.Println(usageMsg)
		return
	}

	// Check for invalid flag formats in the raw args
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--output=") {
			// catches: -output=, --color=, etc.
			fmt.Println(usageMsg)
			return
		}
		if strings.HasPrefix(arg, "output=") || strings.HasPrefix(arg, ".output=") {
			// catches: output=file.txt, .output=file.txt
			fmt.Println(usageMsg)
			return
		}
	}

	// Fix 2 & 4: clearly separate input string from banner style
	var inputStr, bannerStyle string
	if len(arguments) == 1 {
		inputStr = arguments[0]
		bannerStyle = "standard"
	} else {
		// last argument is banner style, everything before is the input
		inputStr = strings.Join(arguments[:len(arguments)-1], " ")
		bannerStyle = arguments[len(arguments)-1]
	}

	// Fix 4: FileReader only needs the banner style, not the output fileName
	contentRead, readingStatus := asciiart.FileReader(bannerStyle)

	if fileName != "" {
		// --output flag was provided: write rendered art to file
		result := asciiart.FormatPrinter(inputStr, string(contentRead), readingStatus)
		writingErr := os.WriteFile(fileName, []byte(result+"\n"), 0666)
		if writingErr != nil {
			// Fix 7: use %v not %W in log.Fatalf
			log.Fatalf("err!: %v", writingErr)
		}
	} else {
		// Fix 1: single code path for stdout — no fall-through double print
		fmt.Println(asciiart.FormatPrinter(inputStr, string(contentRead), readingStatus))
	}
}