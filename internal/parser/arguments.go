package parser

import (
	"fmt"
	"strings"
)

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
	Renames    map[string]string // key: field name, value: acccessor name.
	RawRenames []string

	// Import path name
	ImportPaths    []ImportPath
	RawImportPaths []string
}

// arguments
var Args = Arguments{
	Renames:     map[string]string{},
	ImportPaths: make([]ImportPath, 0),
}

type (
	ImportPath struct {
		Path  string
		Alias string
	}
	ImportPaths []ImportPath
)

// Display import paths
func (is ImportPaths) Display() string {
	if len(is) == 0 {
		return ""
	}

	var txt string
	for _, elem := range is {
		switch elem.Alias {
		case "": // no alias
			txt += fmt.Sprintf("	\"%s\"\n", elem.Path)
		default:
			txt += fmt.Sprintf("	%s \"%s\"\n", elem.Alias, elem.Path)
		}
	}
	return "\nimport (\n" + txt + ")\n"
}

// Load arguments
func (a *Arguments) Load() error {
	errs := make([]error, 0)
	if err := a.loadRename(Args.RawRenames); err != nil {
		errs = append(errs, fmt.Errorf("load rename error: %w", err))
	}

	if err := a.loadImportPath(Args.RawImportPaths); err != nil {
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
			Args.ImportPaths = append(Args.ImportPaths, ImportPath{Path: path})
		case 2: // path:alias case
			path, alias := pair[0], pair[1]
			Args.ImportPaths = append(Args.ImportPaths, ImportPath{Path: path, Alias: alias})
		default:
			errs = append(errs, fmt.Errorf("invalid import path: %s", str))
		}
	}
	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
