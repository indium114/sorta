package internal

import (
	"github.com/charmbracelet/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func RunSorter(dir string) {
	log.Info("Starting sort", "directory", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Error("Failed to read directory", "err", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
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

		category := GetCategory(mime)
		newDir := filepath.Join(dir, category)

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

		log.Info("Moved file",
			"file", entry.Name(),
			"category", category,
			"mime", mime,
		)
	}

	log.Info("Done")
}
