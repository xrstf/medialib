package vorbis

import (
	"fmt"
	"os"
	"strings"

	"go.xrstf.de/medialib/vorbis/types"
)

func WriteChapters(filename string, chapters types.Chapters) error {
	encoded, err := encodeChapters(chapters)
	if err != nil {
		return fmt.Errorf("failed to encode chapters: %w", err)
	}

	return os.WriteFile(filename, []byte(encoded), 0644)
}

// https://wiki.xiph.org/Chapter_Extension
func encodeChapters(chapters types.Chapters) (string, error) {
	var buf strings.Builder

	for i, chap := range chapters {
		num := fmt.Sprintf("%03d", i)

		buf.WriteString(fmt.Sprintf("CHAPTER%s=%s\n", num, chap.Start.UTC().Format("15:04:05.000")))
		buf.WriteString(fmt.Sprintf("CHAPTER%sNAME=%s\n", num, chap.Name))
	}

	return buf.String(), nil
}
