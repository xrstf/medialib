package mkvtoolnix

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/pixelbender/go-matroska/matroska"
)

const mkvinfoBinary = "mkvinfo"

var mkvinfoRegex = regexp.MustCompile(`\| \+ Duration: ([0-9:.]+)`)

func GetDuration(filename string) (time.Duration, error) {
	var stdout bytes.Buffer

	// update mkv file
	cmd := exec.Command(mkvinfoBinary, filename, "--ui-language", "en-us")
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		// mkvinfo hates single quotes on Windows when run through Go:
		// (MKVInfo) Error: Couldn't open source file S03E15 - Yesterday's Enterprise.mkv (open file error).
		// (running mkvinfo manually on the command line works as expected, for some reason)
		// So we just fallback to the slower, Go based parser in these cases.
		dur, err := GetDurationSlow(filename)
		if err != nil {
			return 0, fmt.Errorf("%s failed and so did go-matroska in reading the file: %w", mkvinfoBinary, err)
		}

		return dur, nil
	}

	output := stdout.String()
	match := mkvinfoRegex.FindStringSubmatch(output)

	if match == nil {
		return 0, errors.New("regex didn't match mkvinfo output")
	}

	format := "15:04:05.000000000"

	parsedTime, err := time.ParseInLocation(format, match[1], time.UTC)
	if err != nil {
		return 0, fmt.Errorf("invalid duration: %w", err)
	}

	hours := parsedTime.Hour()
	minutes := parsedTime.Minute()
	seconds := parsedTime.Second()
	nanos := parsedTime.Nanosecond()

	dur := time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(nanos)*time.Nanosecond

	return dur, nil
}

func GetDurationSlow(mkvFile string) (time.Duration, error) {
	doc, err := matroska.Decode(mkvFile)
	if err != nil {
		return 0, err
	}

	// .Duration is milliseconds already, to keep precision we multiply by a thousand
	micros := int64(doc.Segment.Info[0].Duration * 1000)

	return time.Duration(micros) * time.Microsecond, nil
}
