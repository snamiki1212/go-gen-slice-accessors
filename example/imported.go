package main

import "time"

//go:generate go run ../. --entity Imported --slice Importeds --input imported.go --output imported_gen.go
type Imported struct {
	CreatedAt time.Time
}

type Importeds []Imported
