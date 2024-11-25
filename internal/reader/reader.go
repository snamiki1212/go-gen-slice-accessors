package reader

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Reader struct {
	path string
}

func NewReader(path string) *Reader {
	return &Reader{path: path}
}

// Read source code from file.
func (r Reader) Read() (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, r.path, nil, parser.AllErrors)
}
