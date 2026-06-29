package internal

import (
	"github.com/charmbracelet/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// get the category for a given file
func GetCategory(mime string, filename string) string {
	// first pass, based on MIME type
	for category, data := range Categories {
		for _, t := range data.MimeTypes {
			if mime == t {
				return category
			}
		}
	}

	// do a second pass based on the file extension
	ext := strings.ToLower(filepath.Ext(filename))

	for category, data := range Categories {
		for _, e := range data.Extensions {
			if ext == e {
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

func RunSorter(dir string, dry bool) {
	log.Info("Starting sort", "directory", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Error("Failed to read directory", "err", err)
		return
	}

	for _, entry := range entries {
		// skip directories
		if entry.IsDir() {
			continue
		}

		// skip hidden files
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		oldPath := filepath.Join(dir, entry.Name())

		mime, err := DetectMime(oldPath)
		if err != nil {
			log.Error("Failed to detect MIME",
				"file", entry.Name(),
				"err", err,
			)
			continue
		}

		category := GetCategory(mime, entry.Name())
		newDir := filepath.Join(dir, category)

		if !dry {
			if err := os.MkdirAll(newDir, 0755); err != nil {
				log.Error("Failed to create directory",
					"dir", newDir,
					"err", err,
				)
				continue
			}

			newPath := filepath.Join(newDir, entry.Name())

			// Handle collisions
			if _, err := os.Stat(newPath); err == nil {
				newPath = filepath.Join(newDir, "copy_"+entry.Name())
			}

			if err := os.Rename(oldPath, newPath); err != nil {
				log.Error("Failed to move file",
					"file", entry.Name(),
					"err", err,
				)
				continue
			}
		}

		log.Info("Moved file",
			"file", entry.Name(),
			"category", category,
			"mime", mime,
		)
	}

	log.Info("Done")
}
