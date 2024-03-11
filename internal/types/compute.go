package types

type Plugin struct {
	Method       string         `json:"method"`
	Path         string         `json:"path"`
	GlobalConfig map[string]any `json:"global-config"`
}

type Initialize struct {
	Plugins map[string]Plugin `json:"plugins"`
}
type DataModelInterface[T any] interface {
	GetDuration() int
}
type DataModel[T any] struct {
	Duration int
}

type DataChild struct {
	Pipeline []string       `json:"pipeline"`
	Config   any            `json:"config"`
	Defaults map[string]any `json:"defaults"`
	Inputs   []any          `json:"inputs"`
}

type Tree struct {
	Children map[string]DataChild `json:"children"`
}

type Input struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tags        any        `json:"tags"`
	Initialize  Initialize `json:"initialize"`
	Tree        Tree       `json:"tree"`
}
