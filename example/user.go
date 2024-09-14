package main

//go:generate go run ../. --entity User --slice Users --input user.go --output user_gen.go --exclude=CreatedAt,UpdatedAt
type User struct {
	// Example
	UserID string

	// Primitive
	Int          int
	IntPtr       *int
	IntSlice     []int
	IntPtrSlice  []*int
	Bool         bool
	BoolPtr      *bool
	BoolSlice    []bool
	BoolPtrSlice []*bool
	Str          string
	StrPtr       *string
	StrSlice     []string
	StrPtrSlice  []*string

	// Function
	Fn1 func() string
	// FnPtr1 *func() string
	Fn2 func()
	// FnPtr2 *func()
	Fn3 func(func() string) func() int
	// FnPtr3 func(*func() string) *func() int

	// Struct
	// InlineStruct0 struct{ Name string } // NOTE: not supported
	// InlineStruct1 struct { Name string } // NOTE: not supported
	Struct0 DefinedStruct0
	Struct1 DefinedStruct1

	// Map
	Map1         map[string]int
	MapPtr1      *map[string]int
	MapSlice1    []map[string]int
	MapPtrSlice1 []*map[string]int
	Map2         map[string]func() string
	MapPtr2      *map[string]func() string
	MapSlice2    []map[string]func() string
	MapPtrSlice2 []*map[string]func() string

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
