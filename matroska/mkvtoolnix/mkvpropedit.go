package mkvtoolnix

import (
	"bytes"
	"fmt"
	"os/exec"
)

const mkvpropeditBinary = "mkvpropedit"

func SetChapters(mkvFile string, chapterFile string) error {
	if _, err := exec.LookPath(mkvpropeditBinary); err != nil {
		return fmt.Errorf("MKVPropEdit not available: %w", err)
	}

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command(mkvpropeditBinary, mkvFile, "--chapters", chapterFile)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mkvpropedit failed: %s: %w", stdout.String(), err)
	}

	return nil
}
