package main

import "time"

//go:generate go run ../. --entity Imported1 --slice Imported1s --input imported1.go --output imported1_gen.go --import=time
type Imported1 struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Imported1s []Imported1
