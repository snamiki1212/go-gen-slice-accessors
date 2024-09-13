package main

//go:generate go run ../. --entity User --slice Users --input user.go --output user_gen.go --exclude=CreatedAt,UpdatedAt
type User struct {
	// Example
	UserID string

	// Value
	Int     int
	IntPtr  *int
	Bool    bool
	BoolPtr *bool
	Str     string
	StrPtr  *string

	// Function
	// Fn1 func() string
	// PtrFn   *func() string
	// Fn2 func()
	// Fn3 func(func() string) func() int

	// Struct
	// InlineStruct0 struct{}
	// InlineStruct1 struct {
	// 	Name string
	// }
	Struct0 DefinedStruct0
	Struct1 DefinedStruct1

	// Map
	Map1    map[string]int
	MapPtr1 *map[string]int
	Map2    map[string]func() string
	MapPtr2 *map[string]func() string

	// Slices
	// SliInt  []int
	// SliStr  []string
	// SliBool []bool

	// Exclude
	CreatedAt int64
	UpdatedAt int64

	// Channel
	Chan0        chan int
	ChanPtr0     *chan int
	Chan1        chan func() string
	ChanPtr1     *chan func() string
	ChanSend0    chan<- int
	ChanSendPtr0 *chan<- int
	ChanRecv0    <-chan int
	ChanRecvPtr0 *<-chan int
}
type Users []User

type DefinedStruct0 struct{}

type DefinedStruct1 struct {
	Name string
}
