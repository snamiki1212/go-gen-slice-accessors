package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Reader func(path string) (*ast.File, error)

// Read source code from file.
func Read(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, path, nil, parser.AllErrors)
}
