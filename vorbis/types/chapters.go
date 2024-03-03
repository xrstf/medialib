package types

import "time"

type Chapters []Chapter

type Chapter struct {
	Start time.Time
	Name  string
}
