package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parser(t *testing.T) {
	type args struct {
		src       string
		arguments arguments
	}
	tests := map[string]struct {
		args    args
		want    data
		wantErr bool
	}{
		"ok: common": {
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
					{Accessor: "Ages", Name: "Age", Type: "int64"},
				},
			},
		},
		"ok: exlucde": {
			args: args{
				arguments: arguments{entity: "User", slice: "Users", fieldNamesToExclude: []string{"Age"}},
				src: `
package user

type User struct {
	UserID string
	Age    int64
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "UserIDs", Name: "UserID", Type: "string"},
				},
			},
		},
		"ok: custom accessor naming": {
			args: args{
				arguments: arguments{entity: "User", slice: "Users", fieldNamesToExclude: []string{"Age2"}, accessors: map[string]string{"Age": "AgeList", "Age2": "Age2List"}},
				src: `
package user

type User struct {
	UserID string
	Age    int64
	Age2   int64
}
`,
			},
			want: data{
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
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	UserID string
	History bool
	Box bool
}
`,
			},
			want: data{
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
				arguments: arguments{entity: "User", slice: "Users"},
				src: `
package user

type User struct {
	callback0 func()
	callback1 func(x string, x2 bool) (y int64, y2 int32)
	callback2 func(string, bool) (int64, int32)
	callback3 func(u1 User) (u2 *User)
	callback4 func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))
	callback5 func(head string, tail ...bool) (num int64)
}
`,
			},
			want: data{
				pkgName:   "user",
				sliceName: "Users",
				fields: fields{
					{Accessor: "callback0s", Name: "callback0", Type: "func() ()"},
					{Accessor: "callback1s", Name: "callback1", Type: "func(x string, x2 bool) (y int64, y2 int32)"},
					{Accessor: "callback2s", Name: "callback2", Type: "func(string, bool) (int64, int32)"},
					{Accessor: "callback3s", Name: "callback3", Type: "func(u1 User) (u2 *User)"},
					{Accessor: "callback4s", Name: "callback4", Type: "func(cb1 func(x1 string) (y1 int)) (cb2 func(x2 string) (y2 int))"},
					{Accessor: "callback5s", Name: "callback5", Type: "func(head string, tail ...bool) (num int64)"},
				},
			},
		},
		"ok: map": {
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
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
			want: data{
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
		"ng: invalid src code: syntax error": {
			args: args{
				arguments: arguments{entity: "User", slice: "Users"},
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
				arguments: arguments{entity: "User", slice: "Users"},
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
				arguments: arguments{entity: "INVALID_ENTITY", slice: "Users"},
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
			if name != "ok: map" {
				t.Skip()
			}
			reader := newReaderFromString(tt.args.src)
			got, err := parse(tt.args.arguments, reader)
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
func newReaderFromString(src string) func(path string) (*ast.File, error) {
	return func(path string) (*ast.File, error) {
		fset := token.NewFileSet()
		noFilePath := "" // not import from file path
		return parser.ParseFile(fset, noFilePath, src, parser.AllErrors)
	}
}
