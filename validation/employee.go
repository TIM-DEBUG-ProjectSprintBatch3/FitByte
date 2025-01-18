package validation

import (
	"fmt"
)

func isImageURIExtensionValid(uri string) bool {
	extension := uri[len(uri)-5:]
	if extension != ".jpeg" {
		extension = extension[1:]
		if extension != ".jpg" && extension != ".png" {
			fmt.Println("false", extension)
			return false
		}
	}

	return true
}
