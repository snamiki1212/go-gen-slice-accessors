package main

import (
	"bytes"
	// "time"
	alias_time "time"
)

//go:generate go run ../. --entity Imported2 --slice Imported2s --input imported2.go --output imported2_gen.go --import=alias_time:time,bytes
type Imported2 struct {
	Buf       bytes.Buffer
	CreatedAt alias_time.Time
	// UpdatedAt *time.Time
}

type Imported2s []Imported2
