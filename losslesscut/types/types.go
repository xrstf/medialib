package types

type Project struct {
	Version       int          `yaml:"version"`
	MediaFileName string       `yaml:"mediaFileName"`
	CutSegments   []CutSegment `yaml:"cutSegments"`
}

type CutSegment struct {
	Start float64 `yaml:"start"`
	End   float64 `yaml:"end"`
	Name  string  `yaml:"name"`
}
