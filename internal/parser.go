package internal

import (
	"fmt"
	"go/ast"
	"log"
	"slices"
	"strings"

	"github.com/snamiki1212/go-gen-slice-accessors/internal/generator"
)

type Parser struct {
	reader     IReader
	pluralizer IPluralizer
}

type IPluralizer interface {
	Pluralize(str string) string
}

type IReader interface {
	Read() (*ast.File, error)
}

func NewParser(reader IReader, pluralizer IPluralizer) *Parser {
	return &Parser{reader: reader, pluralizer: pluralizer}
}

// Parse
//
// Parse sorce code and new generator.
func (p Parser) Parse(args Arguments) (generator.Generator, error) {
	// Convert source code to ast
	file, err := p.reader.Read()
	if err != nil {
		return generator.Generator{}, fmt.Errorf("parse: error: %w", err)
	}

	// Parse ast
	fields, err := parseFile(file, args)
	if err != nil {
		return generator.Generator{}, err
	}

	// Convert ast to generator struct and transform data
	fs := newFields(fields).ExcludeByFieldName(args.FieldNamesToExclude)
	fs = buildAccessorMulti(fs, p.pluralizer, args.Renames)

	// Parse paths
	paths := func() ImportPaths {
		if args.HasImportPath() {
			return args.ImportPaths
		}
		paths := newImportPathsFromFile(file)
		return filterByUsed(paths, fs)
	}()

	return generator.Generator{
		Fields:      fs,
		PkgName:     getPackageNameFromFile(file),
		SliceName:   args.Slice,
		ImportBlock: paths.Display(),
	}, nil
}

// Get package name.
func getPackageNameFromFile(node *ast.File) string { return node.Name.Name }

// New import paths from file.
func newImportPathsFromFile(node *ast.File) ImportPaths {
	var paths []ImportPath
	for _, imp := range node.Imports {
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		paths = append(paths, ImportPath{Path: strings.Trim(imp.Path.Value, `"`), Alias: alias})
	}
	return paths
}

// Filter import paths by using fields.
func filterByUsed(candidates []ImportPath, fs generator.Fields) []ImportPath {
	ts := []string{} // Actucally Used type name from filed type ex) time.Time -> time
	for _, f := range fs {
		if strings.Contains(f.Type, ".") {
			ts = append(ts, strings.Split(f.Type, ".")[0])
		}
	}

	var res []ImportPath
	for _, tn := range ts {
		for _, imp := range candidates {
			switch imp.Alias {
			case "": // no alias
				path := imp.Path
				// xxx/yyyy/zzz -> zzz
				if strings.Contains(path, "/") {
					path = path[strings.LastIndex(path, "/")+1:]
				}
				if tn == path {
					res = append(res, imp)
				}
			default: // has alias
				if tn == imp.Alias {
					res = append(res, imp)
				}
			}
		}
	}

	// Uniq
	res = slices.CompactFunc(res, func(e1, e2 ImportPath) bool {
		return e1.Path == e2.Path && e1.Alias == e2.Alias
	})

	return res
}

// Parse file.
func parseFile(node *ast.File, args Arguments) ([]*ast.Field, error) {
	// Find entity object
	obj, ok := node.Scope.Objects[args.Entity]
	if !ok {
		return nil, fmt.Errorf("parseFile: entity not found: %s", args.Entity)
	}

	// Find entity
	entity, ok := obj.Decl.(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("parseFile: invalid entity: %s", args.Entity)
	}

	// Find fields
	str, ok := entity.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("parseFile: invalid type: %T", str)
	}
	fs := str.Fields.List

	return fs, nil
}

// Parse expression.
func parseExpr(x ast.Expr) string {
	switch tt := x.(type) {
	case *ast.Ident:
		return parseIdent(tt)
	case *ast.StarExpr:
		return parseStarExpr(tt)
	case *ast.FuncType:
		return parseFuncType(tt)
	case *ast.Ellipsis:
		return parseEllipsis(tt)
	case *ast.MapType:
		return parseMapType(tt)
	case *ast.ChanType:
		return parseChanType(tt)
	case *ast.ArrayType:
		return parseArrayType(tt)
	case *ast.SelectorExpr:
		return parseSelectorExpr(tt)
	}
	log.Println("parseExpr: parse error: unknown type", x)
	return "any"
}

// Parse selector expression.
func parseSelectorExpr(x *ast.SelectorExpr) string {
	return fmt.Sprintf("%s.%s", parseExpr(x.X), parseIdent(x.Sel))
}

// Parse array type.
func parseArrayType(x *ast.ArrayType) string {
	return fmt.Sprintf("[]%s", parseExpr(x.Elt))
}

// Parse chan type.
func parseChanType(x *ast.ChanType) string {
	switch x.Dir {
	case ast.SEND:
		return "chan<- " + parseExpr(x.Value)
	case ast.RECV:
		return "<-chan " + parseExpr(x.Value)
	default:
		return "chan " + parseExpr(x.Value)
	}
}

// Parse map type.
func parseMapType(x *ast.MapType) string {
	return fmt.Sprintf("map[%s]%s", parseExpr(x.Key), parseExpr(x.Value))
}

// Parse identifier.
func parseIdent(x *ast.Ident) string {
	return x.Name
}

// Parse star expression.
func parseStarExpr(x *ast.StarExpr) string {
	switch tt := x.X.(type) {
	case *ast.Ident:
		return "*" + tt.Name
	case *ast.MapType:
		return "*" + parseMapType(tt)
	case *ast.ChanType:
		return "*" + parseChanType(tt)
	case *ast.ArrayType:
		return "*" + parseArrayType(tt)
	case *ast.FuncType:
		return "*" + parseFuncType(tt)
	case *ast.SelectorExpr:
		return "*" + parseSelectorExpr(tt)
	default:
		log.Println("parseStarExpr: parse error: unknown type", x)
		return "any"
	}
}

// Parse function type.
func parseFuncType(x *ast.FuncType) string {
	params := func() string {
		if x != nil && x.Params != nil && x.Params.List != nil {
			return newFields(x.Params.List).Display()
		}
		return ""
	}()
	results := func() string {
		if x != nil && x.Results != nil && x.Results.List != nil {
			return newFields(x.Results.List).Display()
		}
		return ""
	}()
	return fmt.Sprintf("func(%s) (%s)", params, results)
}

// Parse ellipsis.
func parseEllipsis(x *ast.Ellipsis) string {
	str := parseExpr(x.Elt)
	switch str {
	case "any":
		log.Println("parseEllipsis: parse error: unknown type")
		return str
	default:
		return "..." + str
	}
}

// Constructor for fields.
func newFields(raws []*ast.Field) generator.Fields {
	fs := make(generator.Fields, 0, len(raws))
	for _, raw := range raws {
		fs = append(fs, newField(raw))
	}
	return fs
}

// Get field name safely.
func safeGetNameFromField(raw *ast.Field) string {
	if len(raw.Names) == 0 {
		return ""
	}
	return raw.Names[0].Name
}

// ----------------
// Field
// ----------------

// Constructor for field.
func newField(raw *ast.Field) generator.Field {
	name := safeGetNameFromField(raw)
	ty := parseExpr(raw.Type)
	return generator.Field{
		Name: name,
		Type: ty,
	}
}

// Build accessor name.
func buildAccessor(f *generator.Field, p IPluralizer, rule map[string]string) *generator.Field {
	if ac, ok := rule[f.Name]; ok {
		f.Accessor = ac
		return f
	}

	f.Accessor = p.Pluralize(f.Name)
	return f
}

// Build accessor names.
func buildAccessorMulti(fs generator.Fields, p IPluralizer, rule map[string]string) generator.Fields {
	for i := range fs {
		buildAccessor(&fs[i], p, rule)
	}
	return fs
}
