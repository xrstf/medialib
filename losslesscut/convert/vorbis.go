package convert

import (
	"fmt"

	llctypes "go.xrstf.de/medialib/losslesscut/types"
	vorbistypes "go.xrstf.de/medialib/vorbis/types"
)

func ToVorbis(project *llctypes.Project) (vorbistypes.Chapters, error) {
	var chapters vorbistypes.Chapters

	chapterNum := 1

	for _, cut := range project.CutSegments {
		start := secondsToTime(cut.Start)

		name := cut.Name
		if name == "" {
			name = fmt.Sprintf("Chapter %d", chapterNum)
			chapterNum++
		}

		chapters = append(chapters, vorbistypes.Chapter{
			Start: start,
			Name:  name,
		})
	}

	return chapters, nil
}
