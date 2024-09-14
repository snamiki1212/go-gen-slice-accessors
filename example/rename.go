package main

//go:generate go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors --entity Account --slice Accounts --input rename.go --output rename_gen.go --accessor=AccountID:GetAccountIDs --accessor=Age:AgeList
type Account struct {
	AccountID string
	Age       int64
}

type Accounts []Account
