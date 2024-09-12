package main

//go:generate go run ../. --entity User --slice Users --input user.go --output user_gen.go --exclude=CreatedAt,UpdatedAt
type User struct {
	UserID    string
	CreatedAt int64
	UpdatedAt int64
	Chan0     chan int
	Chan1     chan func() string
	ChanA     *chan int
	ChanB     *chan func() string
	ChanX     chan<- int
	ChanY     *chan<- int
}
type Users []User
