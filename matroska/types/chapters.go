package types

type Chapters struct {
	EditionEntry EditionEntry `xml:"EditionEntry"`
}

type EditionEntry struct {
	EditionFlagDefault int           `xml:"EditionFlagDefault"`
	EditionFlagHidden  int           `xml:"EditionFlagHidden"`
	ChapterAtom        []ChapterAtom `xml:"ChapterAtom"`
}

type ChapterAtom struct {
	ChapterUID       uint64         `xml:"ChapterUID"`
	ChapterTimeStart string         `xml:"ChapterTimeStart"`
	ChapterTimeEnd   string         `xml:"ChapterTimeEnd"`
	ChapterDisplay   ChapterDisplay `xml:"ChapterDisplay"`
}

type ChapterDisplay struct {
	ChapterString   string `xml:"ChapterString"`
	ChapterLanguage string `xml:"ChapterLanguage"`
}
