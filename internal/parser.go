package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"slices"
	"strings"
)

// Parse sorce code to own struct.
func Parse(args Arguments, reader func(path string) (*ast.File, error)) (Generator, error) {
	// Convert source code to ast
	file, err := reader(args.Input)
	if err != nil {
		return Generator{}, fmt.Errorf("parse: error: %w", err)
	}

	// Parse ast
	fields, err := parseFile(file, args)
	if err != nil {
		return Generator{}, err
	}

	// Convert ast to own struct
	fs := newFields(fields)

	// Transform data
	fs = fs.
		excludeByFieldName(args.FieldNamesToExclude).
		buildAccessor(newPluralizer(), args.Renames)

	// Parse paths
	paths := func() ImportPaths {
		if args.HasImportPath() {
			return args.ImportPaths
		}
		paths := newImportPathsFromFile(file)
		return filterByUsed(paths, fs)
	}()

	return Generator{
		fields:      fs,
		pkgName:     getPackageNameFromFile(file),
		sliceName:   args.Slice,
		importPaths: paths,
	}, nil
}

// Read source code from file.
func Reader(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, path, nil, parser.AllErrors)
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
		paths = append(paths, ImportPath{path: strings.Trim(imp.Path.Value, `"`), alias: alias})
	}
	return paths
}

// Filter import paths by using fields.
func filterByUsed(candidates []ImportPath, fs fields) []ImportPath {
	ts := []string{} // Actucally Used type name from filed type ex) time.Time -> time
	for _, f := range fs {
		if strings.Contains(f.Type, ".") {
			ts = append(ts, strings.Split(f.Type, ".")[0])
		}
	}

	var res []ImportPath
	for _, tn := range ts {
		for _, imp := range candidates {
			switch imp.alias {
			case "": // no alias
				path := imp.path
				// xxx/yyyy/zzz -> zzz
				if strings.Contains(path, "/") {
					path = path[strings.LastIndex(path, "/")+1:]
				}
				if tn == path {
					res = append(res, imp)
				}
			default: // has alias
				if tn == imp.alias {
					res = append(res, imp)
				}
			}
		}
	}

	// Uniq
	res = slices.CompactFunc(res, func(e1, e2 ImportPath) bool {
		return e1.path == e2.path && e1.alias == e2.alias
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

type (
	fields []field

	// Struct field from entity in source code.
	field struct {
		Name     string // field name like UserID
		Type     string // field type like string,int64...
		Accessor string // accessor name like UserIDs
	}
)

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
			return newFields(x.Params.List).display()
		}
		return ""
	}()
	results := func() string {
		if x != nil && x.Results != nil && x.Results.List != nil {
			return newFields(x.Results.List).display()
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
func newFields(raws []*ast.Field) fields {
	fs := make(fields, 0, len(raws))
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
func newField(raw *ast.Field) field {
	name := safeGetNameFromField(raw)
	ty := parseExpr(raw.Type)
	return field{
		Name: name,
		Type: ty,
	}
}

// Display fields.
func (fs fields) display() string {
	if len(fs) == 0 {
		return ""
	}
	var pairs []string
	for _, f := range fs {
		pairs = append(pairs, f.display())
	}
	return strings.Join(pairs, ", ")
}

// Display field.
func (f field) display() string {
	if f.Name == "" {
		return f.Type
	}
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

// Build accessor name.
func (f *field) buildAccessor(p pluralizer, rule map[string]string) *field {
	if ac, ok := rule[f.Name]; ok {
		f.Accessor = ac
		return f
	}

	f.Accessor = p.pluralize(f.Name)
	return f
}

// Build accessor names.
func (fs fields) buildAccessor(p pluralizer, rule map[string]string) fields {
	for i := range fs {
		fs[i].buildAccessor(p, rule)
	}
	return fs
}

// Exclude fields by name.
func (fs fields) excludeByFieldName(targets []string) fields {
	return slices.DeleteFunc(fs, func(f field) bool {
		return slices.Contains(targets, f.Name)
	})
}
