# go-gen-slice-accessors

Generate accessors for each field in the slice struct.

## Usage

### 1. Add go:generate directive.

```diff filename="user.go"
package main

+//go:generate go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go
type User struct {
  UserID    string
}

type Users []User
```

### 2. Run `go generate` and got generated code.

```diff filename="user_gen.go"
+// Code generated by go generate DO NOT EDIT.
+
+package main
+
+// UserIDs
+func (xs Users) UserIDs() []string {
+	sli := make([]string, 0, len(xs))
+	for i := range xs {
+		sli = append(sli, xs[i].UserID)
+	}
+	return sli
+}
```

## Help

```shell
$ go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors@latest --help

Generate accessors for each field in the slice struct.

Usage:
  gen-slice-accessors [flags]

Flags:
  -a, --accessor strings   accessor name for field / e.g. --accessor=Name:GetName
  -e, --entity string      target entity name
  -x, --exclude strings    field names to exclude
  -h, --help               help for gen-slice-accessors
  -i, --input string       input file name
  -o, --output string      output file name
  -s, --slice string       target slice name
```

## [Examples](./example)

- Exclude fields ([src](./example/exclude.go) / [generated code](./example/exclude_gen.go))
- Rename accessors ([src](./example/rename.go) / [generated code](./example/rename_gen.go))

## E2E

```shell
$ go generate ./example
$ go run ./example
```

## TODO

- Add version name in generated code.

## LICENSE

[MIT](./LICENSE)
