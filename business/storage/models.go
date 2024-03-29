// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package storage

import ()

type Auth struct {
	ID     int64
	ApiKey string
}

type User struct {
	ID       int64
	Username string
}

type UserDatum struct {
	UserID int64
	School string
}

type UserProfile struct {
	UserID    int64
	FirstName string
	LastName  string
	Phone     string
	Address   string
	City      string
}
