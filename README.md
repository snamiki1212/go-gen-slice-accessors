# go-gen-slice-accessors

Generate accessors for each field in the slice struct.

- <b>Easy Setup & Removal</b>
- <b>Code Generation & Zero Package Size</b>

<img src="https://github.com/user-attachments/assets/92602519-44ab-49ad-9093-46fe3858eed3" />

## Usage

### 1. Install `go-gen-sllice-accessors`

```zsh
$ go install github.com/snamiki1212/go-gen-slice-accessors@latest
$ go-gen-slice-accessors --help
# -> To ensure it was installed correctly, otherwise set up your GOPATH like `export PATH=$PATH:$(go env GOPATH)/bin`
```

### 2. Add `go:generate` directive.

```diff filename="user.go"
package main

+//go:generate go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go
type User struct {
  UserID    string
}

type Users []User
```

### 3. Run `go generate` command.

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

### 4. Use accessors

```go
package main

import "fmt"

func main() {
	us := Users{{UserID: "1"}, {UserID: "2"}, {UserID: "3"}}
	ids := us.UserIDs() // 🚀 You can use accessors for slice.
	fmt.Println(ids) // [1 2 3]
}
```

> [!TIP]
> Install a binary using `go:generate` and your team will not need to think about the installation but simly run `go generate`.
>
> ```diff
> +//go:generate go install github.com/snamiki1212/go-gen-slice-accessors@latest
> +//go:generate go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go
>   type User struct {
>     ...
> ```

## Help

```shell
Generate accessors for each field in the slice struct.

Usage:
  gen-slice-accessors [flags]

Flags:
  -e, --entity string     target entity name
  -x, --exclude strings   field names to exclude
  -h, --help              help for gen-slice-accessors
  -m, --import strings    import path name
                           e.g. --import=time
                           e.g. --import=aliasTime:time
  -i, --input string      input file name
  -o, --output string     output file name
  -r, --rename strings    rename accessor name
                           e.g. --rename=Name:GetName
  -s, --slice string      target slice name
```

## Examples

Generated example Codes.

- [Common case](./example/user_gen.go) ([source](./example/user.go))
- [Exclude flag](./example/exclude_gen.go) ([source](./example/exclude.go))
- [Rename flag](./example/rename_gen.go) ([source](./example/rename.go))
- [Private field case](./example/private_gen.go) ([source](./example/private.go))

## Contribution

### E2E

```shell
$ go generate ./example
$ go run ./example
```

## Alternatives

If you like this package and would like to generate more, please refer to the following.

- [go-gen-lo](https://github.com/snamiki1212/go-gen-lo): Generate samber/lo method for your struct.

## LICENSE

[MIT](./LICENSE)
