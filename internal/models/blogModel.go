package models

import "time"

type Post struct {
	Title    string
	Text     string
	DateTime time.Time
	Author User
	Like []Like
}

type Like struct {
	User User
}

type Komment struct {
	User User
	Text string
	DateTime time.Time
}