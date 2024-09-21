package main

//go:generate go run ../. --entity Private --slice Privates --input private.go --output private_gen.go
type Private struct {
	Public  string
	private string
}

type Privates []Private
