package internal

import (
	"io"
	"net/http"
	"os"
)

// get the category for a given MIME type
func GetCategory(mime string) string {
	for category, types := range Categories {
		for _, t := range types {
			if mime == t {
				return category
			}
		}
	}
	return "Other"
}

// detects the MIME type of a file using its content
func DetectMime(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}
