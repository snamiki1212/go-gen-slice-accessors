package cmd

import (
	"fmt"
	"strings"
)

type Arguments struct {
	// Target entity name
	entity string

	// Target slice name
	slice string

	// Input file name
	input string

	// Output file name
	output string

	// Field names to exclude
	fieldNamesToExclude []string

	// Mapping field name to renamed name
	renames map[string]string // key: field name, value: acccessor name.

	// Import path name
	importPaths []importPath
}

// Load arguments
func (a *Arguments) load() error {
	errs := make([]error, 0)
	if err := a.loadRename(renames); err != nil {
		errs = append(errs, fmt.Errorf("load rename error: %w", err))
	}

	if err := a.loadImportPath(importPaths); err != nil {
		errs = append(errs, fmt.Errorf("load import path error: %w", err))
	}

	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}

	return nil
}

// HasImportPath
func (a *Arguments) HasImportPath() bool {
	return len(a.importPaths) != 0
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
		args.renames[field] = rename
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
			args.importPaths = append(args.importPaths, importPath{path: path})
		case 2: // path:alias case
			path, alias := pair[0], pair[1]
			args.importPaths = append(args.importPaths, importPath{path: path, alias: alias})
		default:
			errs = append(errs, fmt.Errorf("invalid import path: %s", str))
		}
	}
	if len(errs) != 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
