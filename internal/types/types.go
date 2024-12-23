package types

type Place struct {
	Name        string
	Shortcut    string
	Description string
	Path        string
	DocsUrl     string `yaml:"docs-url"`
	Hidden      bool   `yaml:"hidden" default:"false"`
}

func (p *Place) ShortcutAsRune() rune {
	if p.Shortcut == "" {
		return '_'
	}
	r := []rune(p.Shortcut)
	return r[0]
}
