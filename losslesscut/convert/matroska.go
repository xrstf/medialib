package convert

import (
	"fmt"
	"time"

	llctypes "go.xrstf.de/medialib/losslesscut/types"
	mkvtypes "go.xrstf.de/medialib/matroska/types"
)

func ToMatroska(project *llctypes.Project, mediaDuration time.Duration) (*mkvtypes.Chapters, error) {
	var chapters []mkvtypes.ChapterAtom

	finalEnd := durationToTimestamp(mediaDuration)
	chapterNum := 1

	for idx, cut := range project.CutSegments {
		start := secondsToTimestamp(cut.Start)

		end := finalEnd
		if cut.End != 0 {
			end = secondsToTimestamp(cut.End)
		}

		name := cut.Name
		if name == "" {
			name = fmt.Sprintf("Chapter %d", chapterNum)
			chapterNum++
		}

		chapters = append(chapters, mkvtypes.ChapterAtom{
			ChapterUID:       uint64(idx + 1),
			ChapterTimeStart: start,
			ChapterTimeEnd:   end,
			ChapterDisplay: mkvtypes.ChapterDisplay{
				ChapterString:   name,
				ChapterLanguage: "eng",
			},
		})
	}

	return &mkvtypes.Chapters{
		EditionEntry: mkvtypes.EditionEntry{
			EditionFlagDefault: 1,
			EditionFlagHidden:  0,
			ChapterAtom:        chapters,
		},
	}, nil
}
