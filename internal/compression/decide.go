package compression

import (
	"path/filepath"
	"strings"
)

func ShouldCompress(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".zip", ".gz", ".rar", ".7z", ".mp4", ".mkv", ".jpg", ".jpeg", ".png", ".pdf":
		return false
	default:
		return true
	}
}