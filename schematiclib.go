package gokicadlib

import "bytes"

type SchematicLibrary struct {
	Symbols []Symbol
}

func (sl *SchematicLibrary) AddSymbol(ss ...Symbol) {
	for _, s := range ss {
		sl.Symbols = append(sl.Symbols, s)
	}
}

func (sl *SchematicLibrary) KicadLib() *bytes.Buffer {
	var output bytes.Buffer

	output.WriteString("EESchema-LIBRARY Version 2.3\r\n#encoding utf-8\r\n")

	for _, s := range sl.Symbols {
		output.WriteString(s.KicadLib())
	}

	output.WriteString("#\r\n#End Library\r\n")

	return &output
}
