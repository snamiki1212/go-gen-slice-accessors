package main

import (
	"bytes"
	alias_bytes "bytes"
	"time"
)

//go:generate go run ../. --entity Imported --slice Importeds --input imported.go --output imported_gen.go
type Imported struct {
	Buf1      bytes.Buffer
	AliasBuf1 alias_bytes.Buffer
	ValMonth1 time.Month
	ValMonth2 time.Month // duplicated but should be uniqed import.
	PtrTime1  *time.Time
	PtrTime2  *time.Time // duplicated but should be uniqed import.
}

type Importeds []Imported
