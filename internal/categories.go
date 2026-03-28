package internal

var Categories = map[string][]string{
	"Archives": {
		"application/zip",
		"application/x-tar",
		"application/gzip",
		"application/x-7z-compressed",
		"appliction/vnd.rar",
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
		"audio/midi",
	},
	"Video": {
		"video/mp4",
		"video/x-matroska",
	},
	"Documents": {
		"application/pdf",
		"text/plain",
		"text/markdown",
		"application/vnd.oasis.opendocument.presentation",
		"application/vnd.oasis.opendocument.spreadsheet",
		"application/vnd.oasis.opendocument.text",
	},
	"Code": {
		"text/javascript",
		"application/json",
		"application/x-httpd-php",
	},
}
