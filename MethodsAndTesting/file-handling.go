package asciiart

import (
	"fmt"
	"os"
)

func FileHandler(fileName, styleName string) ([]byte, bool) {
	if styleName != "standard" && styleName != "thinkertoy" && styleName != "shadow" {
		styleName = "standard"
	}
	// Read banner file
	data, err := os.ReadFile("banners/" + styleName + ".txt")
	if err != nil {
		fmt.Println("Error")
		return []byte{}, false // returns an empty slice of byte and false if there's is an error
	}


	return data, true
}