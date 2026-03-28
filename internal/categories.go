package internal

var Categories = map[string][]string{
	"Archives": {
		"application/zip",
		"application/x-tar",
		"application/gzip",
		"application/x-7z-compressed",
	},
	"Images": {
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
	},
	"Audio": {
		"audio/mpeg",
		"audio/wav",
		"audio/ogg",
	},
	"Video": {
		"video/mp4",
		"video/x-matroska",
	},
	"Documents": {
		"application/pdf",
		"text/plain",
	},
}
