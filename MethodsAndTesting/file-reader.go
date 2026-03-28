package asciiart

import (
	"fmt"
	"os"
)

func FileReader(styleName string) ([]byte, bool) {
	if styleName != "standard" && styleName != "thinkertoy" && styleName != "shadow" {
		styleName = "standard"
	}

	data, readingErr := os.ReadFile("banners/" + styleName + ".txt")
	if readingErr != nil {
		fmt.Println("Error")
		return []byte{}, false
	}

	return data, true
}