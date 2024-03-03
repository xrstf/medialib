package matroska

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"go.xrstf.de/medialib/matroska/types"
)

func WriteChapters(filename string, chapters *types.Chapters) error {
	encoded, err := encodeChapters(chapters)
	if err != nil {
		return fmt.Errorf("failed to encode chapters: %w", err)
	}

	return os.WriteFile(filename, encoded, 0644)
}

// https://www.matroska.org/technical/chapters.html
func encodeChapters(chapters *types.Chapters) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("<?xml version=\"1.0\"?>\n")
	buf.WriteString("<!-- <!DOCTYPE Chapters SYSTEM \"matroskachapters.dtd\"> -->\n")

	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	if err := encoder.Encode(chapters); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
