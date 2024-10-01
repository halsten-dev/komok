package filepath

import (
	"path/filepath"
	"regexp"
	"strings"
)

var windowsFileNameRegex, _ = regexp.Compile(`[\\/:*?<>|]`)

func MakeWindowSystemFriendlyName(n string) string {
	return windowsFileNameRegex.ReplaceAllString(n, "_")
}

func ChangeExtension(path, newExt string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext) + newExt
}

func AssureExtension(path, ext string) string {
	extension := filepath.Ext(path)

	if extension != ext {
		return ChangeExtension(path, ext)
	}

	return path
}
