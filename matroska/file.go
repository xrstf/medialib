package matroska

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"go.xrstf.de/medialib/matroska/types"
)

func WriteChapters(xmlFile string, chapters *types.Chapters) error {
	encoded, err := encodeChapters(chapters)
	if err != nil {
		return fmt.Errorf("failed to encode chapters: %w", err)
	}

	// safety check against catastrophic bugs
	if !strings.HasSuffix(xmlFile, ".xml") {
		panic(fmt.Sprintf("refusing to write to non-XML filename %q (did you mean to use mkvtoolnix.SetChapters?)", xmlFile))
	}

	return os.WriteFile(xmlFile, encoded, 0644)
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
