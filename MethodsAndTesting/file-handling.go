package asciiart

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FileHandler(paras ...string) ([]byte, bool) {
	var mainStr strings.Builder
	var fileName strings.Builder
	var styleName strings.Builder
	if len(paras) == 2 {
		fileName.WriteString(paras[0])
		styleName.WriteString(paras[1])
	} else if len(paras) == 1 {
		mainStr.WriteString(paras[0])
	}

	if styleName.String() != "standard" && styleName.String() != "thinkertoy" && styleName.String() != "shadow" {
		styleName.WriteString("standard")
	}
	// Read banner file
	data,readingErr := os.ReadFile("banners/" + styleName.String() + ".txt")
	if readingErr != nil {
		fmt.Println("Error")
		return []byte{}, false // returns an empty slice of byte and false if there's is an error
	}

	writingErr := os.WriteFile(fileName.String(), []byte(FormatPrinter(mainStr.String(), "", true)), 0666)	
	if writingErr != nil {
		log.Fatalf("err!: %W", writingErr)
	}	

	return data, true
}