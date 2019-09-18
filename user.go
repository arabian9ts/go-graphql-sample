package main

type User struct {
	ID       int64
	Name     string
}

type Users []*User

type Message struct {
	ID     int64
	UserID int64
	Text   string
}

type Messages []*Message
