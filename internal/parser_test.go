package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/snamiki1212/go-gen-slice-accessors/internal/pluralizer"
	"github.com/stretchr/testify/assert"
)

func Test_parser(t *testing.T) {
	type args struct {
		src       string
		arguments Arguments
	}
	tests := map[string]struct {
		args    args
		want    Generator
		wantErr bool
	}{
		"ok: common": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
					{Accessor: "Ages", Name: "Age", Type: "int64"},
				},
			},
		},
		"ok: imported": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

import "time"

type User struct {
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	UpdatedAt2 *time.Time
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "CreatedAts", Name: "CreatedAt", Type: "time.Time"},
					{Accessor: "UpdatedAts", Name: "UpdatedAt", Type: "*time.Time"},
					{Accessor: "UpdatedAt2s", Name: "UpdatedAt2", Type: "*time.Time"},
				},
				importBlock: (ImportPaths{
					{path: "time", alias: ""},
				}).Display(),
			},
		},
		"ok: exlucde": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users", FieldNamesToExclude: []string{"Age"}},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
				},
			},
		},
		"ok: rename": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users", FieldNamesToExclude: []string{"Age2"}, Renames: map[string]string{"Age": "AgeList", "Age2": "Age2List"}},
				src: `
package user

type User struct {
	UserID string
	Age    int64
	Age2   int64
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
					{Accessor: "AgeList", Name: "Age", Type: "int64"},
				},
			},
		},
		"ok: plural": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	UserID string
	History bool
	Box bool
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
					{Accessor: "Histories", Name: "History", Type: "bool"},
					{Accessor: "Boxes", Name: "Box", Type: "bool"},
				},
			},
		},
		"ok: callback": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	callback0 func()
	callback1 func(x string, x2 bool) (y int64, y2 int32)
	callback2 func(string, bool) (int64, int32)
	callback3 func(u1 User) (u2 *User)
	callback4 func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))
	callback5 func(head string, tail ...bool) (num int64)
	callback6 *func(x string) (y int64)
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "callback0s", Name: "callback0", Type: "func() ()"},
					{Accessor: "callback1s", Name: "callback1", Type: "func(x string, x2 bool) (y int64, y2 int32)"},
					{Accessor: "callback2s", Name: "callback2", Type: "func(string, bool) (int64, int32)"},
					{Accessor: "callback3s", Name: "callback3", Type: "func(u1 User) (u2 *User)"},
					{Accessor: "callback4s", Name: "callback4", Type: "func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))"},
					{Accessor: "callback5s", Name: "callback5", Type: "func(head string, tail ...bool) (num int64)"},
					{Accessor: "callback6s", Name: "callback6", Type: "*func(x string) (y int64)"},
				},
			},
		},
		"ok: map": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	map0 map[string]string
	map1 map[string]func()
	mapA *map[string]string
	mapB *map[string]func()
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "map0s", Name: "map0", Type: "map[string]string"},
					{Accessor: "map1s", Name: "map1", Type: "map[string]func() ()"},
					{Accessor: "mapAs", Name: "mapA", Type: "*map[string]string"},
					{Accessor: "mapBs", Name: "mapB", Type: "*map[string]func() ()"},
				},
			},
		},
		"ok: slice": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	SliceInt []int
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "SliceInts", Name: "SliceInt", Type: "[]int"},
				},
			},
		},
		"ok: chan": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	chan0 chan string
	chan1 chan func()
	chanA *chan string
	chanB *chan func()
	chanS0 chan<- string
	chanR0 <-chan string
}
`,
			},
			want: Generator{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "chan0s", Name: "chan0", Type: "chan string"},
					{Accessor: "chan1s", Name: "chan1", Type: "chan func() ()"},
					{Accessor: "chanAs", Name: "chanA", Type: "*chan string"},
					{Accessor: "chanBs", Name: "chanB", Type: "*chan func() ()"},
					{Accessor: "chanS0s", Name: "chanS0", Type: "chan<- string"},
					{Accessor: "chanR0s", Name: "chanR0", Type: "<-chan string"},
				},
			},
		},
		"ng: invalid src code: syntax error": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
package user

type User struct {
	UserID string
}
hogehoge // syntax error
`,
			},
			wantErr: true,
		},
		"ng: invalid src code: not found package name": {
			args: args{
				arguments: Arguments{Entity: "User", Slice: "Users"},
				src: `
// no package name
type User struct {
	UserID string
}
`,
			},
			wantErr: true,
		},
		"ng: invalid arguments: not found entity": {
			args: args{
				arguments: Arguments{Entity: "INVALID_ENTITY", Slice: "Users"},
				src: `
package user
type User struct {
	UserID string
}
`,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			reader := NewReaderFromString(tt.args.src)
			parser := NewParser(reader, pluralizer.NewPluralizer())
			got, err := parser.Parse(tt.args.arguments)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Construct reader from string.
type ReaderFromString struct {
	src string
}

func NewReaderFromString(src string) ReaderFromString {
	return ReaderFromString{src: src}
}

func (r ReaderFromString) Read() (*ast.File, error) {
	fset := token.NewFileSet()
	noFilePath := "" // not import from file path
	return parser.ParseFile(fset, noFilePath, r.src, parser.AllErrors)
}
