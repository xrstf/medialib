package edl

import (
	"fmt"
	"os"
	"strings"
)

func WriteFile(edlFile string, mediaFile string, chapters []int) error {
	encoded := encodeChapters(mediaFile, chapters)

	// safety check against catastrophic bugs
	if !strings.HasSuffix(edlFile, ".edl") {
		panic(fmt.Sprintf("refusing to write to non-edl filename %q", edlFile))
	}

	return os.WriteFile(edlFile, []byte(encoded), 0644)
}

// https://github.com/mpv-player/mpv/blob/master/DOCS/edl-mpv.rst
func encodeChapters(mediaFile string, chapters []int) string {
	edlLines := []string{"# mpv EDL v0"}
	mediaLen := len(mediaFile)

	// this generates a "dumb" file, where instead of "chapters N to M", we just list
	// all chapters to play individually
	for _, chapterID := range chapters {
		// to allow filenames to have commas, we must use the special "%length%data" syntax
		edlLines = append(
			edlLines,
			fmt.Sprintf("%%%d%%%s,%d,1,timestamps=chapters", mediaLen, mediaFile, chapterID),
		)
	}

	return strings.Join(edlLines, "\n")
}
