package util

import (
	"path/filepath"
	"strings"
)

func ReplaceExtension(filename string, newExt string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + newExt
}
