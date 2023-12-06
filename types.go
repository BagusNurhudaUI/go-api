package main

import "math/rand"

type Account struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func newAccount(fn string, ln string) *Account {
	return &Account{
		ID:        rand.Int63n(1000),
		FirstName: fn,
		LastName:  ln,
		Number:    rand.Int63n(100000),
		Balance:   rand.Int63n(100000),
	}
}
