package main

//go:generate go run ../. --entity User --slice Users --input user.go --output user_gen.go --exclude=CreatedAt,UpdatedAt
type User struct {
	// Example
	UserID string

	// Value
	Int  int
	Bool bool
	Str  string

	// Pointer
	PtrInt  *int
	PtrStr  *string
	PtrBool *bool
	// PtrFn   *func() string

	// Function
	// Fn1 func() string
	// Fn2 func()
	// Fn3 func(func() string) func() int

	// Struct
	// Item0 struct{}
	// Item1 struct {
	// 	Name string
	// }

	// Map
	Map1 map[string]int
	Map2 map[string]func() string

	// Slices
	// SliInt  []int
	// SliStr  []string
	// SliBool []bool

	// Exclude
	CreatedAt int64
	UpdatedAt int64

	// Channel
	Chan0 chan int
	Chan1 chan func() string
	ChanA *chan int
	ChanB *chan func() string
	ChanX chan<- int
	ChanY *chan<- int
}
type Users []User
