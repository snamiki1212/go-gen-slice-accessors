package cmd

import (
	"bytes"
	"fmt"
	"text/template"
)

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
func generate(data data, args arguments) (string, error) {
	pkgName := data.pkgName
	sliceName := data.sliceName
	infos := data.fields

	if len(infos) == 0 {
		return "", nil
	}

	var txt string

	// append header
	txt += "// Code generated by \"go-gen-slice-accessors\"; DO NOT EDIT.\n"
	txt += "// Based on information from https://github.com/snamiki1212/go-gen-slice-accessors\n"
	txt += "\n"
	txt += fmt.Sprintf("package %s\n", pkgName)
	txt += func() string {
		var paths []importPath
		if args.HasImportPath() {
			paths = args.importPaths
		} else {
			paths = data.importPaths
		}
		return GenerateImportPath(paths)
	}()

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

	return txt, nil
}
