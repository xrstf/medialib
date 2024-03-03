package mkvtoolnix

import (
	"time"

	"github.com/pixelbender/go-matroska/matroska"
)

func GetDuration(mkvFile string) (time.Duration, error) {
	doc, err := matroska.Decode(mkvFile)
	if err != nil {
		return 0, err
	}

	// .Duration is milliseconds already, to keep precision we multiply by a thousand
	micros := int64(doc.Segment.Info[0].Duration * 1000)

	return time.Duration(micros) * time.Microsecond, nil
}
