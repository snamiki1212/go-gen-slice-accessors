package internal

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"
)

// Generator: code generation struct
type Generator struct {
	fields      fields
	pkgName     string
	sliceName   string
	importBlock string
}

const templateBody = `
// {{ .Method }}
func (xs {{ .Slice }}) {{ .Method }}() []{{ .Type }} {
	sli := make([]{{ .Type }}, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].{{ .Field }})
	}
	return sli
}
`

// Replace variable from key to value in template.
type templateMapper struct {
	Slice  string // Slice name for target struct (ex. Users).
	Method string // Method name of accessor (ex. UserIDs).
	Type   string // Type name of field (ex. string).
	Field  string // Field name of struct (ex. UserID).
}

// Generate code
func (g Generator) Generate() (string, error) {
	pkgName := g.pkgName
	sliceName := g.sliceName
	infos := g.fields

	if len(infos) == 0 {
		return "", nil
	}

	var txt string

	// append header
	txt += "// Code generated by \"go-gen-slice-accessors\"; DO NOT EDIT.\n"
	txt += "// Based on information from https://github.com/snamiki1212/go-gen-slice-accessors\n"
	txt += "\n"
	txt += fmt.Sprintf("package %s\n", pkgName)
	txt += g.importBlock

	// append templates
	var doc bytes.Buffer
	tp, err := template.New("").Parse(templateBody)
	if err != nil {
		return "", fmt.Errorf("template parse error: %w", err)
	}
	for _, info := range infos {
		data := &templateMapper{
			Slice:  sliceName,
			Method: info.Accessor,
			Type:   info.Type,
			Field:  info.Name,
		}

		err = tp.Execute(&doc, data)
		if err != nil {
			return "", fmt.Errorf("template execute error: %w", err)
		}
	}
	txt += doc.String()

	// format (go fmt)
	btxt, err := format.Source([]byte(txt))
	if err != nil {
		return "", fmt.Errorf("format error: %w", err)
	}

	txt = string(btxt)

	return txt, nil
}
