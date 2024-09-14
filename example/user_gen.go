// Code generated by "go-gen-slice-accessors"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-slice-accessors

package main

// UserIDs
func (xs Users) UserIDs() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].UserID)
	}
	return sli
}

// Ints
func (xs Users) Ints() []int {
	sli := make([]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Int)
	}
	return sli
}

// IntPtrs
func (xs Users) IntPtrs() []*int {
	sli := make([]*int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].IntPtr)
	}
	return sli
}

// Bools
func (xs Users) Bools() []bool {
	sli := make([]bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Bool)
	}
	return sli
}

// BoolPtrs
func (xs Users) BoolPtrs() []*bool {
	sli := make([]*bool, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].BoolPtr)
	}
	return sli
}

// Strs
func (xs Users) Strs() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Str)
	}
	return sli
}

// StrPtrs
func (xs Users) StrPtrs() []*string {
	sli := make([]*string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].StrPtr)
	}
	return sli
}

// Fn1s
func (xs Users) Fn1s() []func() (string) {
	sli := make([]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn1)
	}
	return sli
}

// Fn2s
func (xs Users) Fn2s() []func() () {
	sli := make([]func() (), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn2)
	}
	return sli
}

// Fn3s
func (xs Users) Fn3s() []func(func() (string)) (func() (int)) {
	sli := make([]func(func() (string)) (func() (int)), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Fn3)
	}
	return sli
}

// Struct0s
func (xs Users) Struct0s() []DefinedStruct0 {
	sli := make([]DefinedStruct0, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Struct0)
	}
	return sli
}

// Struct1s
func (xs Users) Struct1s() []DefinedStruct1 {
	sli := make([]DefinedStruct1, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Struct1)
	}
	return sli
}

// Map1s
func (xs Users) Map1s() []map[string]int {
	sli := make([]map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Map1)
	}
	return sli
}

// MapPtr1s
func (xs Users) MapPtr1s() []*map[string]int {
	sli := make([]*map[string]int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapPtr1)
	}
	return sli
}

// Map2s
func (xs Users) Map2s() []map[string]func() (string) {
	sli := make([]map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Map2)
	}
	return sli
}

// MapPtr2s
func (xs Users) MapPtr2s() []*map[string]func() (string) {
	sli := make([]*map[string]func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].MapPtr2)
	}
	return sli
}

// Chan0s
func (xs Users) Chan0s() []chan int {
	sli := make([]chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Chan0)
	}
	return sli
}

// ChanPtr0s
func (xs Users) ChanPtr0s() []*chan int {
	sli := make([]*chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanPtr0)
	}
	return sli
}

// Chan1s
func (xs Users) Chan1s() []chan func() (string) {
	sli := make([]chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Chan1)
	}
	return sli
}

// ChanPtr1s
func (xs Users) ChanPtr1s() []*chan func() (string) {
	sli := make([]*chan func() (string), 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanPtr1)
	}
	return sli
}

// ChanSendPtr0s
func (xs Users) ChanSendPtr0s() []*chan<- int {
	sli := make([]*chan<- int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanSendPtr0)
	}
	return sli
}

// ChanRecv0s
func (xs Users) ChanRecv0s() []<-chan int {
	sli := make([]<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecv0)
	}
	return sli
}

// ChanRecvPtr0s
func (xs Users) ChanRecvPtr0s() []*<-chan int {
	sli := make([]*<-chan int, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].ChanRecvPtr0)
	}
	return sli
}
