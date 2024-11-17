package internal

import (
	"fmt"
	"strings"
)

// arguments
var Args = Arguments{
	Renames:     map[string]string{},
	ImportPaths: make([]ImportPath, 0),
}

// Import path name
var RawImportPaths []string
var RawRenames []string

type ImportPath struct {
	path  string
	alias string
}

// GenerateImportPath
func GenerateImportPath(importPaths []ImportPath) string {
	if len(importPaths) == 0 {
		return ""
	}

	var txt string
	for _, elem := range importPaths {
		switch elem.alias {
		case "": // no alias
			txt += fmt.Sprintf("	\"%s\"\n", elem.path)
		default:
			txt += fmt.Sprintf("	%s \"%s\"\n", elem.alias, elem.path)
		}
	}
	return "\nimport (\n" + txt + ")\n"
}

type Arguments struct {
	// Target Entity name
	Entity string

	// Target Slice name
	Slice string

	// Input file name
	Input string

	// Output file name
	Output string

	// Field names to exclude
	FieldNamesToExclude []string

	// Mapping field name to renamed name
	Renames map[string]string // key: field name, value: acccessor name.

	// Import path name
	ImportPaths []ImportPath
}

// Load arguments
func (a *Arguments) Load() error {
	errs := make([]error, 0)
	if err := a.loadRename(RawRenames); err != nil {
		errs = append(errs, fmt.Errorf("load rename error: %w", err))
	}

	if err := a.loadImportPath(RawImportPaths); err != nil {
		errs = append(errs, fmt.Errorf("load import path error: %w", err))
	}

	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}

	return nil
}

// HasImportPath
func (a *Arguments) HasImportPath() bool {
	return len(a.ImportPaths) != 0
}

func (a *Arguments) loadRename(as []string) error {
	errs := make([]error, 0)
	for _, ac := range as {
		pair := strings.Split(ac, ":")
		if len(pair) != 2 {
			errs = append(errs, fmt.Errorf("invalid rename: %s", ac))
			continue
		}
		field, rename := pair[0], pair[1]
		Args.Renames[field] = rename
	}
	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

func (a *Arguments) loadImportPath(sli []string) error {
	errs := make([]error, 0)
	for _, str := range sli {
		pair := strings.Split(str, ":")
		switch len(pair) {
		case 1: // only path case
			path := pair[0]
			Args.ImportPaths = append(Args.ImportPaths, ImportPath{path: path})
		case 2: // path:alias case
			path, alias := pair[0], pair[1]
			Args.ImportPaths = append(Args.ImportPaths, ImportPath{path: path, alias: alias})
		default:
			errs = append(errs, fmt.Errorf("invalid import path: %s", str))
		}
	}
	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
