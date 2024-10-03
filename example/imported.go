package main

import (
	"bytes"
	"time"
	alias_time "time"
)

//go:generate go run ../. --entity Imported --slice Importeds --input imported.go --output imported_gen.go --import=bytes,time:alias_time,time
type Imported struct {
	Buf       bytes.Buffer
	CreatedAt alias_time.Time
	UpdatedAt *time.Time
}

type Importeds []Imported
