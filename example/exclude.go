package main

//go:generate go run ../. --entity Item --slice Items --input exclude.go --output exclude_gen.go --exclude=CreatedAt,UpdatedAt
type Item struct {
	ItemID    string
	CreatedAt int64
	UpdatedAt int64
}

type Items []Item
