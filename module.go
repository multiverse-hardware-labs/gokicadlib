package gokicadlib

import "fmt"

// Module Kicad Module
type Module struct {
	Name string

	Layer        Layer
	Tedit        TimeStamp
	Tstamp       TimeStamp
	Origin       Point
	Descr        string
	Attributes   []string
	Tags         []string
	Path         string
	Reference    Text
	Value        Text
	Arcs         []Arc
	Lines        []Line
	Pads         []Pad
	Model        Model
	Text         []Text
	SExpressions []SExpression
}

// ToSExp convert to kicad s-expression
func (m Module) ToSExp() string {
	module := JoinStrings("module", m.Reference.Text, NewSExp("layer", false, string(m.Layer)), NewSExp("tedit", false, m.Tedit.Stamp().String()), NewSExp("tstamp", false, m.Tstamp.String()))
	descr := NewSExp("descr", false, fmt.Sprintf("\"%s\"", m.Descr))
	attr := NewSExp("attr", false, m.Attributes...)
	at := fmt.Sprintf("(at %.4f %.4f)", m.Origin.X, m.Origin.Y)
	tags := NewSExp("tags", false, m.Tags...)
	reference := m.Reference.ToSExp()
	value := m.Value.ToSExp()
	lines := LineSlice(m.Lines).ToSExp()
	arcs := ArcSlice(m.Arcs).ToSExp()
	pads := PadSlice(m.Pads).ToSExp()
	text := TextSlice(m.Text).ToSExp()
	ses := SExpressionSlice(m.SExpressions).ToSExp()
	return NewSExp(module, true, descr, attr, at, tags, reference, value, lines, arcs, pads, text, ses)
}
